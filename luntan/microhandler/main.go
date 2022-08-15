package main

import (
	"luntan/microhandler/handler"
	pb "luntan/microhandler/proto"

	"github.com/asim/go-micro/v3"

	"github.com/micro/micro/v3/service/logger"

	//consul 1
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/registry"
)

// protoc --proto_path=. --micro_out=. --go_out=. *.proto

const (
	ServerName = "consul-test"
	ConsulAddr = "192.168.226.128:8500" //集群的节点？
)

func main() {
	// Create service
	// 	// consul 3
	// 	// consul注册
	consulReg := consul.NewRegistry(
		registry.Addrs(ConsulAddr),
	)
	srv := micro.NewService(
		micro.Name(ServerName),
		micro.Registry(consulReg),
	)

	// Register handlerr
	pb.RegisterHelloworldHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

// /修改 3
// 	// consul 3
// 	// consul注册
// 	consulReg := consul.NewRegistry(
// 		registry.Addrs(ConsulAddr),
// 	)
// 	srv := service.NewService( // service.New
// 		service.Name(ServerName),    // 服务名字
// 		service.Registry(consulReg), // 注册中心
// 	)

// 	// Register handler
// 	//修改 4
// 	//注册我们的服务
// 	_ = pb.RegisterTestHandler(srv.Server(), new(handler.Test))

// 	// Run service
// 	if err := srv.Run(); err != nil {
// 		logger.Fatal(err)
// 	}
