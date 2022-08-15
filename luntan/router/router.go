package router

import (
	"luntan/middleware"

	"github.com/gin-gonic/gin"
)

//想做一个修改信息的功能 网上查查怎么发邮箱验证
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.Static("/home/gua/go/src/luntan/front/layui", "front/layui") //后替换 前

	r.LoadHTMLGlob("html/*")
	r.GET("/registerhtml", func(c *gin.Context) {
		c.HTML(200, "register.html", gin.H{})
	})
	r.GET("/loginhtml", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{})
	})
	r.GET("/luntan", func(c *gin.Context) {
		c.HTML(200, "luntan.html", gin.H{})
	})
	r.GET("/tiezi", func(c *gin.Context) {
		c.HTML(200, "tiezi.html", gin.H{})
	})
	r.GET("/owner", func(c *gin.Context) {
		c.HTML(200, "owner.html", gin.H{})
	})
	r.GET("/paihang", func(c *gin.Context) {
		c.HTML(200, "paihangbang.html", gin.H{})
	})
	r.GET("/guanzhulist", func(c *gin.Context) {
		c.HTML(200, "guanzhulist.html", gin.H{})
	})
	r.POST("/register", CreatUserInDenglu)
	r.GET("/login", GetUserinDenglu)
	user := r.Group("/logined")
	user.Use(middleware.JWT())
	{
		user.GET("paihang", Userpaihang)
		user.GET("guanzhuget", Guanzhuget)
		user.GET("userget", UserGet)
		user.POST("userupdate", UserUpdate)
		user.POST("guanzhu", Guanzhu)
		user.GET("guanzhulist", GuanzhuList)
		user.POST("dianzan", Tiezidianzan)
		user.GET("dianzanget", Tiezidianget)
		user.GET("tiezixiangqing", TieziGet)
		user.POST("tiezicreate", TieziCreate)
		user.GET("tiezilist", TieziList)
		user.GET("tiezipinglunlist", TieziListPinglun)
		user.POST("tiezipingluncreate", TieziPinglunCreate)
	}

	return r
}
