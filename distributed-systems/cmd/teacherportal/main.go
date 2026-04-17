package main

import (
	"context"
	"fmt"
	stlog "log"

	"git-server.git-server/code-ecosystem/distributed-systems/pkg/log"
	"git-server.git-server/code-ecosystem/distributed-systems/pkg/registry"
	"git-server.git-server/code-ecosystem/distributed-systems/pkg/service"
	"git-server.git-server/code-ecosystem/distributed-systems/pkg/teacherportal"
)

func main() {
	err := teacherportal.ImportTemplates()
	if err != nil {
		stlog.Fatal(err)
	}

	host, port := "localhost", "5000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.TeacherPortal
	r.ServiceUrl = serviceAddress
	r.RequiredServices = []registry.ServiceName{
		registry.LogService,
		registry.GradingService,
	}
	r.ServiceUpdateUrl = r.ServiceUrl + "/services"
	r.HeartBeatURL = serviceAddress + "/heartbeat"

	ctx, err := service.Start(context.Background(), r, host, port, teacherportal.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service is found %v.\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()
	fmt.Println("Shutting down teachers portal service")
}
