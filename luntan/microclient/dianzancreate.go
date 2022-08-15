package microclient

import (
	dianzanpb "luntan/microdianzan/proto"
	"luntan/model"

	"github.com/asim/go-micro/v3"
	"github.com/gin-gonic/gin"

	//consul 1
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/registry"
)

func Dianzancreateclient(param model.Dianzan, c *gin.Context) (*dianzanpb.CreateResponse, error) {
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.188.128:8500"),
	)

	service := micro.NewService(
		micro.Registry(consulReg), //设置注册中心
	)

	service.Init()

	// // 创建微服务客户端client 一个类grpc
	// client := testpb.NewTestService("test", service.Client()) //对应服务端的 name：test服务
	//consul 3 创建微服务客户端【将服务名称从test改为为consul-test】
	client := dianzanpb.NewDianzanService("consul-dianzanservice", service.Client()) //从consul的节点进行读取服务返回
	// 调用服务
	rsp, err := client.Create(c, &dianzanpb.Createequest{ //客户端调用
		Uid:      int64(param.Uid),
		Postid:   int64(param.Postid),
		Followid: int64(param.Followid),
	})
	return rsp, err
}
