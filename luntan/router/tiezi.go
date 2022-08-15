package router

import (
	"fmt"
	"log"
	"luntan/global"
	"luntan/microclient"
	"luntan/model"
	"luntan/pkg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TieziList(c *gin.Context) {
	pageroffSet := c.Query("page")
	page, err := strconv.Atoi(pageroffSet)
	if err != nil {
		log.Println("strconv err")
	}
	// totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	// if err != nil {
	// 	global.Logger.Infof("svc.CountTag err:%v", err)
	// 	response.ToErrorResponse(errcode.ErrorCountTagFail)
	// 	return
	// }

	// tags, err := model.TieziList(page, 10, global.DbEngine)
	tags, err := microclient.Tiezilistclient(page, c)
	if err != nil {
		log.Println("帖子list  err：", err)
		return
	}

	c.JSON(200, gin.H{
		"code":  0, //0表示成功？   在前端 我副了 重点！！！！！！！！！！！！！！！！！！！！！！！！！
		"msg":   "获取成功",
		"data":  tags.List, //List 是_List的slice
		"count": 1000,
	})
	return

}
func TieziListPinglun(c *gin.Context) {
	param := model.NewTieziPinglun()
	valid, errs := pkg.BindAndValid(c, &param)
	fmt.Printf("gg%v", valid)
	if !valid {
		log.Println("帖子list valid err：", errs)
		c.JSON(500, gin.H{
			"code": 200,
			"msg":  "帖子评论获取失败",
		})
		return
	}
	pageroffSet := c.Query("page")
	page, err := strconv.Atoi(pageroffSet)
	if err != nil {
		log.Println("strconv err", err)
	}
	rsp, err := microclient.Tiezipinglunlistclient(param, page, c)
	if err != nil {
		log.Println("帖子评论list  err：", errs)
		c.JSON(500, gin.H{
			"code": 200,
			"msg":  "帖子评论获取失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"tiezipinglunList": rsp.List, //*************
		"code":             0,
		"msg":              "帖子评论获取成功",
		"count":            1000,
	})
	return

}

//详情界面
func TieziGet(c *gin.Context) {
	param := model.NewTiezi()
	valid, errs := pkg.BindAndValid(c, &param)
	fmt.Printf("gg%v", valid)
	if !valid {
		log.Println("帖子list valid err：", errs)
		return
	}
	rsp, err := microclient.Tiezigetclient(param, c)
	if err != nil {
		log.Println("帖子list  err：", errs)
		return
	}
	c.JSON(200, gin.H{
		"postid":   rsp.Id,
		"uid":      rsp.Userid,
		"title":    rsp.Title,
		"content":  rsp.Content,
		"photourl": rsp.Photourl,
	})
	return
}
func TieziCreate(c *gin.Context) {
	param := model.NewTiezi()
	p := UploadFile(c)
	valid, errs := pkg.BindAndValid(c, &param)
	param.Photourl = p
	if !valid {
		fmt.Println("Tiezi Create valid err:", errs)
		//return
	}
	// err := param.TieziCreate(global.DbEngine)
	//tiezicreate服务
	rsp, err := microclient.Tiezicreateclient(param, c)
	if err != nil {
		fmt.Println("Tiezi create err:", err)
		c.JSON(500, gin.H{
			"msg": "创建失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"postid": rsp.Id,
		"mesg":   "创建成功",
	})
	return
}

//在涉及的操作上做gorm事务

func Tiezidianzan(c *gin.Context) {
	param := model.NewDianzan()
	valid, errs := pkg.BindAndValid(c, &param)
	if !valid {
		fmt.Println("dianzan  valid err:", errs)
		//return
	}
	//使用gorm事务 适合电表一次操作
	// tiezidianzanget服务 如果存在 就是取消点赞（删除）
	// userupdate服务user点赞数 + 1
	// 帖子update服务帖子点赞数 + 1
	_, err := microclient.Dianzanclient(param, c)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "点赞失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "点赞成功",
	})
	return

}

func Tiezidianget(c *gin.Context) {
	param := model.NewDianzan()
	valid, errs := pkg.BindAndValid(c, &param)
	if !valid {
		fmt.Println("dianzanget  valid err:", errs)
		//return
	}
	//使用gorm事务 适合电表一次操作
	// tiezidianzanget服务 如果存在 就是取消点赞（删除）
	// userupdate服务user点赞数 + 1
	// 帖子update服务帖子点赞数 + 1
	rsp, err := microclient.Dianzangetclient(param, c)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "点赞获取失败",
		})
		return
	}
	if rsp.Flag == -1 {
		c.JSON(200, gin.H{
			"msg":  "未点赞",
			"code": 0,
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":  "已点赞",
		"code": 0,
	})
	return

}

func TieziPinglunGet(c *gin.Context) {
	param := model.NewTieziPinglun()
	valid, errs := pkg.BindAndValid(c, &param)
	fmt.Printf("gg%v", valid)
	if !valid {
		log.Println("帖子list valid err：", errs)
		return
	}
	tags, err := param.TieziPinglunGet(global.DbEngine)
	if err != nil {
		log.Println("帖子list  err：", errs)
		return
	}
	c.JSON(200, gin.H{
		"postid":        tags.Postid,
		"postcommentid": tags.ID,
		"id":            tags.Uid,
		"time":          tags.Time,
		"content":       tags.Content,
	})
	return
}
func TieziPinglunCreate(c *gin.Context) {
	param := model.NewTieziPinglun()
	valid, errs := pkg.BindAndValid(c, &param)
	if !valid {
		fmt.Println("Tiezipinglun Create valid err:", errs)
		return
	}
	err := param.TieziPinglunCreate(global.DbEngine)
	if err != nil {
		fmt.Println("Tiezipinglun create err:", err)
		return
	}
	c.JSON(200, gin.H{
		"msg": "创建成功",
	})
	return
}
