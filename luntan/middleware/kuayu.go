package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func Cors() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		method := c.Request.Method

// 		origin := c.Request.Header.Get("Origin")

// 		if origin != "" {
// 			c.Header("Access-Control-Allow-Origin", origin)
// 			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization") //自定义 Header
// 			c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT")
// 			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
// 			c.Header("Access-Control-Allow-Credentials", "true")

// 		}

// 		if method == "OPTIONS" {
// 			c.Header("Access-Control-Allow-Origin", "*")
// 			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization") //自定义 Header
// 			c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT")
// 			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
// 			c.Header("Access-Control-Allow-Credentials", "true")
// 			c.AbortWithStatus(http.StatusNoContent)
// 		}

// 		c.Next()
// 	}
// }

// func Cors() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		method := context.Request.Method
// 		context.Header("Access-Control-Allow-Origin", "*")
// 		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
// 		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,PUT")
// 		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
// 		context.Header("Access-Control-Allow-Credentials", "true")
// 		if method == "OPTIONS" {
// 			context.AbortWithStatus(http.StatusNoContent)
// 		}
// 		context.Next()
// 	}
// }
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			// 当Access-Control-Allow-Credentials为true时，将*替换为指定的域名
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-Extra-Header, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Max-Age", "86400") // 可选
			c.Set("content-type", "application/json")   // 可选
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
