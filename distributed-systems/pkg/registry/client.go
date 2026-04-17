package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
)

func RegisterService(r Registration) error {
	heartBeatUrl, err := url.Parse(r.HeartBeatURL)
	if err != nil {
		return err
	}
	http.HandleFunc(heartBeatUrl.Path, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	serviceUpdateUrl, err := url.Parse(r.ServiceUpdateUrl)
	if err != nil {
		return err
	}
	http.Handle(serviceUpdateUrl.Path, &serviceUpdateHandler{})

	buff := new(bytes.Buffer)
	enc := json.NewEncoder(buff)
	err = enc.Encode(r)
	if err != nil {
		return err
	}
	res, err := http.Post(ServiceUrl, "application/json", buff)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register service. Registry service responded with code %v\n", res.StatusCode)
	}
	return nil
}

type serviceUpdateHandler struct{}

func (suh serviceUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	dec := json.NewDecoder(r.Body)
	var p patch
	err := dec.Decode(&p)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Printf("Update received: %+v\n", p)
	prov.Update(p)
}

func ShutdownService(serviceUrl string) error {
	req, err := http.NewRequest(http.MethodDelete, ServiceUrl, bytes.NewBuffer([]byte(serviceUrl)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "text/plain")
	res, err := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to deregister service. Registry service responded with code %v\n", res.StatusCode)
	}
	return nil
}

type providers struct {
	services map[ServiceName][]string
	mutex    *sync.RWMutex
}

func (p *providers) Update(pat patch) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for _, patchEntry := range pat.Added {
		if _, ok := p.services[patchEntry.Name]; !ok {
			p.services[patchEntry.Name] = make([]string, 0)

		}
		p.services[patchEntry.Name] = append(p.services[patchEntry.Name], patchEntry.Url)
	}
	for _, patchEntry := range pat.Removed {
		if providerUrls, ok := p.services[patchEntry.Name]; ok {
			for i := range providerUrls {
				if providerUrls[i] == patchEntry.Url {
					p.services[patchEntry.Name] = append(p.services[patchEntry.Name][:i], providerUrls[i+1:]...)
				}
			}
		}
	}
}

func (p providers) get(name ServiceName) (string, error) {
	providers, ok := p.services[name]
	if !ok {
		return "", fmt.Errorf("No providers available for %v.\n", name)
	}
	idx := int(rand.Float32() * float32(len(providers)))
	return providers[idx], nil
}

func GetProvider(name ServiceName) (string, error) {
	return prov.get(name)
}

var prov = providers{
	services: make(map[ServiceName][]string),
	mutex:    new(sync.RWMutex),
}
