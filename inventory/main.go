package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"syscall"
	"time"

	"git-server.git-server/code-ecosystem/inventory/inventorypb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func main() {
	app := http.Server{
		Addr: ":8080",
	}
	p := inventorypb.Product{
		Id:          int32(inventory[0].Id),
		Name:        inventory[0].Name,
		Description: inventory[0].Description,
		Stock:       int32(inventory[0].Stock),
	}
	data, err := proto.Marshal(&p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	go startGRPCServer()
	time.Sleep(2 * time.Second)
	callGRPCService()

	http.HandleFunc("/products", allProducts)
	http.HandleFunc("/product/", productByid)
	http.HandleFunc("/product/add", productAdd)
	http.HandleFunc("/product/delete/", delProductById)
	http.HandleFunc("/product/update", updateProduct)

	log.Println("Application started at port :8080")
	go func() {
		err := app.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down gracefully")
	app.Shutdown(context.Background())
	log.Println("App stopped gracefully")
}

func allProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	enc := json.NewEncoder(w)
	w.Header().Add("content-type", "application/json")
	enc.SetIndent("", " ")
	enc.Encode(map[string]any{
		"data":   inventory,
		"status": http.StatusOK,
	})
}

func productByid(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	re := regexp.MustCompile(`^\/product\/(\d+?)$`)
	matches := re.FindStringSubmatch(r.URL.Path)
	if len(matches) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(matches[1])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for _, item := range inventory {
		if item.Id == id {
			enc := json.NewEncoder(w)
			w.Header().Add("content-type", "application/json")
			enc.SetIndent("", " ")
			enc.Encode(map[string]any{
				"data":   item,
				"status": http.StatusOK,
			})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func productAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body := r.Body
	defer body.Close()
	var data Product
	dec := json.NewDecoder(body)
	dec.Decode(&data)
	data.Id = inventory.NewId()
	inventory = append(inventory, data)
	w.WriteHeader(http.StatusAccepted)
}

func delProductById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	re := regexp.MustCompile(`^\/product\/delete\/(\d+?)$`)
	matches := re.FindStringSubmatch(r.URL.Path)
	if len(matches) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(matches[1])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for i, item := range inventory {
		if item.Id == id {
			inventory = append(inventory[:i], inventory[i+1:]...)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body := r.Body
	defer body.Close()
	var data Product
	dec := json.NewDecoder(body)
	dec.Decode(&data)
	for j := range inventory {
		if inventory[j].Id == data.Id {
			inventory[j] = data
			w.WriteHeader(http.StatusAccepted)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

type ProductService struct {
	inventorypb.UnimplementedProductServer
}

func (ps ProductService) GetProduct(ctx context.Context, req *inventorypb.GetProductRequest) (*inventorypb.GetProductReply, error) {
	for _, p := range inventory {
		if p.Id == int(req.Id) {
			return &inventorypb.GetProductReply{
				Product: &inventorypb.Product{
					Id:          int32(p.Id),
					Name:        p.Name,
					Description: p.Description,
					Stock:       int32(p.Stock),
				},
			}, nil
		}
	}
	return nil, fmt.Errorf("product not found")
}

func startGRPCServer() {
	lis, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	inventorypb.RegisterProductServer(grpcServer, &ProductService{})
	log.Fatal(grpcServer.Serve(lis))
}

func callGRPCService() {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("localhost:4000", opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := inventorypb.NewProductClient(conn)
	res, err := client.GetProduct(context.TODO(), &inventorypb.GetProductRequest{Id: 2})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", res.Product)

}
