package main

import (
	"context"
	"fmt"
	"os"

	stlog "log"

	"git-server.git-server/code-ecosystem/distributed-systems/pkg/log"
	"git-server.git-server/code-ecosystem/distributed-systems/pkg/registry"
	"git-server.git-server/code-ecosystem/distributed-systems/pkg/service"
)

func main() {
	log.Run("./app.log")
	host, port := os.Getenv("HOST"), os.Getenv("PORT")
	if len(host) == 0 {
		host = "localhost"
	}
	if len(port) == 0 {
		port = "4000"
	}
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.LogService
	r.ServiceUrl = serviceAddress
	r.RequiredServices = make([]registry.ServiceName, 0)
	r.ServiceUpdateUrl = r.ServiceUrl + "/services"
	r.HeartBeatURL = serviceAddress + "/heartbeat"
	ctx, err := service.Start(context.Background(), r, host, port, log.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down log service")
}
