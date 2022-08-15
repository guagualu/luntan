package router

import (
	"log"
	"luntan/microclient"
	"luntan/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Userpaihang(c *gin.Context) {
	//注册
	//c.HTML(http.StatusOK, "index.html", nil)
	param := model.NewUser()
	pageroffSet := c.Query("page")
	page, err := strconv.Atoi(pageroffSet)
	if err != nil {
		log.Println("strconv err", err)
	}
	rsp, err := microclient.Userpaihangclient(param, page, c)
	if err != nil {
		log.Println("user paihang err ", err)
		c.JSON(500, gin.H{
			"msg":  "排行榜失败",
			"code": 200,
			"data": rsp.List,
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":   "排行成功",
		"code":  0,
		"data":  rsp.List,
		"count": 1000,
	})
	return
}
