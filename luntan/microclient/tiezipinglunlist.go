package microclient

import (
	tiezipb "luntan/microtiezi/proto"
	"luntan/model"

	"github.com/asim/go-micro/v3"
	"github.com/gin-gonic/gin"

	//consul 1
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/registry"
)

func Tiezipinglunlistclient(param model.TieziPinglun, page int, c *gin.Context) (*tiezipb.TiezipinglunlistResponse, error) {
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
	rsp, err := client.Listpinglun(c, &tiezipb.Tiezipinglunlistrequest{
		Page:   int64(page),
		Postid: int64(param.Postid),
	})
	return rsp, err
}
