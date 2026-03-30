package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	mu := &sync.RWMutex{}
	cacheChan := make(chan Book)
	dbChan := make(chan Book)
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1

		wg.Add(2)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			defer wg.Done()
			if b, ok := queryCache(id, m); ok {
				ch <- b
			}
		}(id, wg, mu, cacheChan)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			defer wg.Done()
			if b, ok := queryDatabase(id, m); ok {
				ch <- b
			}
		}(id, wg, mu, dbChan)
		go func(cacheChan, dbChan <-chan Book) {
			select {
			case b := <-cacheChan:
				fmt.Println("Got value from cache:")
				fmt.Println(b)
				<-dbChan
			case b := <-dbChan:
				fmt.Println("Got value from database:")
				fmt.Println(b)
			}
		}(cacheChan, dbChan)
		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (Book, bool) {
	time.Sleep(200 * time.Microsecond)
	for _, b := range books {
		if b.Id == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}
	return Book{}, false
}
