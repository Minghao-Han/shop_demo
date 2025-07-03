package main

import (
	"context"
	"example_shop/kitex_gen/example/shop/item"
	"example_shop/kitex_gen/example/shop/item/itemservice"
	tools "example_shop/shared"
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	hregistry "github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/utils"
	consulapi "github.com/hashicorp/consul/api"
	hconsul "github.com/hertz-contrib/registry/consul"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	cli itemservice.Client
)

func main() {
	// Get local ip address
	ip, err := tools.GetLocalIPv4Address()
	if err != nil {
		log.Fatal(err)
	}
	addr := ip + ":8889"

	// Kitex consul resolver
	consul_addr := "172.17.0.2:8500"
	resolver, err := consul.NewConsulResolver(consul_addr)
	if err != nil {
		log.Fatal(err)
	}

	// Discover item service
	c, err := itemservice.NewClient("item.server", client.WithResolver(resolver), client.WithRPCTimeout(time.Second*3))
	if err != nil {
		log.Fatal(err)
	}
	cli = c

	// Register this service to consul
	config := consulapi.DefaultConfig()
	config.Address = consul_addr
	consulClient, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
		return
	}
	// build a consul register with the consul client
	r := hconsul.NewConsulRegister(consulClient)

	hz := server.Default(
		server.WithHostPorts(addr),
		server.WithRegistry(r, &hregistry.Info{
			ServiceName: "api",
			Addr:        utils.NewNetAddr("tcp", addr),
			Weight:      10,
			Tags:        nil,
		}),
	)

	hz.GET("/api/item", Handler)

	// if err := hz.Run(); err != nil {
	// 	log.Fatal(err)
	// }
	hz.Spin()
}

func Handler(ctx context.Context, c *app.RequestContext) {
	req := item.NewGetItemReq()
	req.Id = 1024
	resp, err := cli.GetItem(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	c.String(200, resp.String())
}
