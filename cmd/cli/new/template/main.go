package template

var (
	MainSRV = `package main

import (
	"{{.Dir}}/handler"
	pb "{{.Dir}}/proto"

	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	log "github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"

)

func main() {
	// Create service
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs =[]string{"consul.kanter.cn:80"}
	})
	service := micro.NewService(
		micro.Name("{{lower .Alias}}"),
		micro.Version("latest"),
		micro.Registry(reg),
	)
	// Initialise service
	service.Init()

	// Register handler
	pb.Register{{title .Alias}}Handler(service.Server(), new(handler.{{title .Alias}}))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
`
)
