package microclient

import (
	tiezipb "luntan/microtiezi/proto"

	"github.com/asim/go-micro/v3"
	"github.com/gin-gonic/gin"

	//consul 1
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/registry"
)

func Tiezilistclient(page int, c *gin.Context) (*tiezipb.TiezilistResponse, error) {
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
	client := tiezipb.NewTieziService("consul-tieziservice", service.Client()) //从consul的节点进行读取服务返回
	// 调用服务
	rsp, err := client.List(c, &tiezipb.Tiezilistrequest{
		Pageoffset: int64(page),
	})
	return rsp, err
}
