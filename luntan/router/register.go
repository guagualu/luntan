package router

import (
	"fmt"
	"log"
	"luntan/microclient"
	"luntan/model"
	"luntan/pkg"

	"github.com/gin-gonic/gin"
)

func CreatUserInDenglu(c *gin.Context) {
	//注册
	//c.HTML(http.StatusOK, "index.html", nil)
	param := model.NewUser()
	valid, errs := pkg.BindAndValid(c, &param)
	fmt.Println(param)
	if !valid {
		log.Println("CreatUserInRegister bindandvalid err:", errs)
		return
	}
	p := UploadFile(c)
	param.PhotoUrl = p
	rsp, err := microclient.Registerclient(param, c)
	if err != nil {
		log.Println("register  err:", err)
		c.JSON(500, gin.H{
			"message": "创建失败",
		})
		return
	}
	if rsp.Flag == -1 {
		log.Println("register flag err ")
		c.JSON(500, gin.H{
			"message": "创建失败",
		})
		return
	}
	c.JSON(200, gin.H{ //返回的 可以用来做回调
		"message": "创建成功",
		"id":      rsp.Id,
	})

	return
}
