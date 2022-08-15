package router

import (
	"fmt"
	"log"
	"luntan/microclient"
	"luntan/model"
	"luntan/pkg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserinDenglu(c *gin.Context) {
	param := model.NewUser()
	//response := app.NewResponse(c)
	valid, errs := pkg.BindAndValid(c, &param) //获取与验证
	fmt.Println("test sno duibudui", param)
	if !valid {
		log.Println("bindandvalid errs :%v", errs)
	}
	rsp, err := microclient.Userdenglugetclient(param, c)
	if err != nil {
		log.Println("GEtUser err:", err)
		c.JSON(500, gin.H{
			"message": "登陆失败",
		})
		return
	}
	theid := strconv.Itoa(int(param.ID))
	token, err := pkg.GenerateToken(theid, rsp.Password) //生成token
	if err != nil {
		log.Println(" GenerateToken err:", err)
		c.JSON(500, gin.H{
			"message": "token验证失败",
		})
		return
	}
	// fmt.Println("token is")
	// fmt.Println(token)
	c.JSON(200, gin.H{
		"token":  token, //因为前端那里需要接收一个data*****************
		"id":     rsp.Id,
		"secret": rsp.Password,
	})

}
