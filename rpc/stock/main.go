package main

import (
	stock "example_shop/kitex_gen/example/shop/stock/stockservice"
	tools "example_shop/shared"
	"log"
	"net"
	"strconv"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	ip, err := tools.GetLocalIPv4Address()
	if err != nil {
		log.Fatal(err)
	}
	r, err := consul.NewConsulRegister("172.17.0.2:8500")
	if err != nil {
		log.Fatal(err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", net.JoinHostPort(ip, strconv.Itoa(8500)))
	svr := stock.NewServer(new(StockServiceImpl), server.WithServiceAddr(addr), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "stock.server",
	}))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
