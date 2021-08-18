package template

var (
	MainClientSRV = `package main

import (
	"context"
	"fmt"
    pb "{{.Dir}}/proto"

	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3"
	
)

func main() {
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs =[]string{"consul.kanter.cn:80"}
	})
	service := micro.NewService(
		micro.Registry(reg),
	)
	service.Init()
	client := pb.New{{title .Alias}}Service("{{lower .Alias}}",service.Client())
	resp,_ := client.Call(context.TODO(),&pb.Request{
		Name:"justcy",
	})
	fmt.Println(resp)
}

`
)
