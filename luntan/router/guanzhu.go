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

func Guanzhuget(c *gin.Context) {
	param := model.NewGuanzhu()
	valid, errs := pkg.BindAndValid(c, &param) //获取与验证
	if !valid {
		log.Println("bindandvalid errs :%v", errs)
	}
	// guanzhu服务
	// 先get关注 如果有记录 则delete 如果没有 则create
	rsp, err := microclient.Guanzhugetclient(param, c)
	if err != nil {
		fmt.Println("Guanzhu get err:", err)
		c.JSON(500, gin.H{
			"msg":  "关注获取失败",
			"code": 1,
		})
		return
	}
	if rsp.Flag == 1 {
		c.JSON(200, gin.H{
			"msg":  "已关注",
			"code": 0,
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":  "未关注",
		"code": 0,
	})

	return
}
func Guanzhu(c *gin.Context) {
	param := model.NewGuanzhu()
	valid, errs := pkg.BindAndValid(c, &param) //获取与验证
	if !valid {
		log.Println("bindandvalid errs :%v", errs)
	}
	// guanzhu服务
	// 先get关注 如果有记录 则delete 如果没有 则create
	rsp, err := microclient.Guanzhugetclient(param, c)
	if err != nil {
		fmt.Println("Guanzhu get err:", err)
		c.JSON(500, gin.H{
			"msg": "关注操作失败",
		})
		return
	}
	if rsp.Flag == 1 {
		_, err := microclient.Guanzhudeleteclient(param, c)
		if err != nil {
			c.JSON(500, gin.H{
				"msg": "关注操作失败",
			})
			return
		}
		c.JSON(200, gin.H{
			"msg": "取消关注成功",
		})
		return
	}
	if rsp.Flag == -1 {
		_, err := microclient.Guanzhucreateclient(param, c)
		if err != nil {
			c.JSON(500, gin.H{
				"msg":  "关注操作失败",
				"code": 1,
			})
			return
		}
		c.JSON(200, gin.H{
			"msg":  "关注成功",
			"code": 0,
		})
	}
	return
}
func GuanzhuList(c *gin.Context) {
	suid := c.Query("uid")
	uid, err := strconv.Atoi(suid)

	if err != nil {
		log.Println("uids strconv err", err)
		return
	}
	param := model.NewGuanzhu()
	param.Uid = uid
	pageroffSet := c.Query("page")
	page, err := strconv.Atoi(pageroffSet)
	if err != nil {
		log.Println("strconv err")
	}
	rsp, err := microclient.Guanzhulistclient(param, page, c)
	rellist := make([]model.User, 0)
	for i := 0; i < len(rsp.List); i++ {
		param1 := model.NewUser()
		param1.ID = uint(rsp.List[i].Followid)
		rsp1, err := microclient.Usergetclient(param1, c)
		if err != nil {
			log.Println("user get  err", err)
			c.JSON(500, gin.H{
				"code": 200,
				"msg":  "获取关注列表失败",
			})
			return
		}
		param1.Mail = rsp1.Mail
		param1.Name = rsp1.Name
		param1.Dianzanshu = int(rsp1.Dianzanshu)
		rellist = append(rellist, param1)
	}
	if err != nil {
		fmt.Println("Guanzhu list err:", err)
		c.JSON(500, gin.H{
			"code": 200,
			"msg":  "获取关注列表失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"data":  rellist,
		"code":  0,
		"msg":   "获取关注列表成功",
		"count": 1000,
	})
}
