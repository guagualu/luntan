package pkg

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

//分页功能
//分页处理 page参数为页数 pageSize参数为每页显示数量
// func GetPage(c *gin.Context) int {
// 	page := c.Query("page")//将query参数转为int
// 	if page <= 0 {
// 		return 1
// 	}
// 	return page
// }

// func GetPageSize(c *gin.Context) int { //每页显示的数据数量
// 	PageSize := c.Query("page_size")
// 	if PageSize <= 0 {
// 		return globalself.XiangQingSize //设定的参数
// 	return PageSize
// }

func GetPageOffset(PageSize int, c *gin.Context) int { //要跳过的的数据的个数
	page := c.Query("page") //转换
	Page, err := strconv.Atoi(page)
	if err != nil {
		fmt.Println("page 转换错误")
	}
	result := 0
	if PageSize > 0 {
		result = (Page - 1) * PageSize
	}
	return result
}
