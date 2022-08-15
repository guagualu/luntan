package router

import (
	"log"
	"luntan/microclient"
	"luntan/model"
	"luntan/pkg"

	"github.com/gin-gonic/gin"
)

func UserGet(c *gin.Context) {
	//注册
	//c.HTML(http.StatusOK, "index.html", nil)
	param := model.NewUser()
	valid, errs := pkg.BindAndValid(c, &param)
	if !valid {
		log.Println("userget bindandvalid err:", errs)
		return
	}
	rsp, err := microclient.Usergetclient(param, c)
	if err != nil {
		log.Println("user get err ", err)
		c.JSON(500, gin.H{
			"message": "userget失败",
		})
		return
	}
	user := model.NewUser()
	user.ID = uint(rsp.Id)
	user.Mail = rsp.Mail
	user.Name = rsp.Name
	user.Password = rsp.Password
	user.PhotoUrl = rsp.Photourl
	c.JSON(200, gin.H{
		"message": "get成功",
		"user":    user,
	})
	return
}
func UserdengluGet(c *gin.Context) {
	//注册
	//c.HTML(http.StatusOK, "index.html", nil)
	param := model.NewUser()
	valid, errs := pkg.BindAndValid(c, &param)
	if !valid {
		log.Println("userget bindandvalid err:", errs)
		return
	}
	rsp, err := microclient.Userdenglugetclient(param, c)
	if err != nil {
		log.Println("user denglu get err ", err)
		c.JSON(500, gin.H{
			"message": "userget失败",
		})
		return
	}
	user := model.NewUser()
	user.ID = uint(rsp.Id)
	user.Mail = rsp.Mail
	user.Name = rsp.Name
	user.Password = rsp.Password
	user.PhotoUrl = rsp.Photourl
	c.JSON(200, gin.H{
		"message": "get成功",
		"user":    user,
	})
	return
}

func UserUpdate(c *gin.Context) {
	//注册
	//c.HTML(http.StatusOK, "index.html", nil)
	param := model.NewUser()
	valid, errs := pkg.BindAndValid(c, &param)
	p := UploadFile(c)
	param.PhotoUrl = p
	if !valid {
		log.Println("userget bindandvalid err:", errs)
		return
	}
	rsp, err := microclient.Userupdateclient(param, c)
	if err != nil {
		log.Println("user get err ", err)
		c.JSON(500, gin.H{
			"message": "userupdate失败",
		})
		return
	}
	user := model.NewUser()
	user.ID = uint(rsp.Id)
	user.Mail = rsp.Mail
	user.Name = rsp.Name
	user.Password = rsp.Password
	user.PhotoUrl = rsp.Photourl
	c.JSON(200, gin.H{
		"message": "update成功",
		"user":    user,
	})
	return
}
