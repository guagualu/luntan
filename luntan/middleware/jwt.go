package middleware

import (
	"fmt"
	"luntan/microclient"
	"luntan/model"
	"luntan/pkg"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//作用所有路由的中间件

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			flag  int = 1
		)
		if s, exist := c.GetQuery("Authorization"); exist {
			token = s
		} else {
			token = c.GetHeader("Authorization") //从header中获取token
			// token = token[7:]
		}
		// token = c.Query("token") //从header中获取token
		if token == "" {
			fmt.Println("token is null")
			flag = 0
		} else {
			_, err := pkg.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					fmt.Println("parse token jwt.ValidationErrorExpired")
					flag = 0
				default:
					flag = 0
					fmt.Println("parse token unknow err")
				}
			}
		}
		fmt.Println("jwt里")
		// fmt.Println(token)
		claims, err := pkg.ParseToken(token) //解析发回来的header里面的token 返回需要的信息
		if err != nil {
			fmt.Printf("解析出错%v", err)
			// flag = 0
		}
		u := model.NewUser()
		nowid, _ := strconv.Atoi(claims.Id)
		u.ID = uint(nowid)
		u.Password = claims.Secret
		//getuser服务
		_, err = microclient.Usergetclient(u, c)
		if err != nil {
			fmt.Printf("获取user错误%v", err)
			// return
		}

		//db.Where("app_key=?", userID).Find(&user) //从数据库中找数据 用jwt的db操作封装函数
		// c.Set("userid", User.Id) //可以在请求上下文里面设置一些值，然后其他地方取值 c.Get("user")  //取值 实现了跨中间件取值
		if flag != 1 {
			c.JSON(500, gin.H{
				"message": "鉴权失败",
			})
			c.Abort()
			return
		}
		c.Next() //执行后面的中间件
	}
}
