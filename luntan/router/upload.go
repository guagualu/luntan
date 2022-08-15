package router

import (
	"fmt"
	"luntan/service"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) string {
	fmt.Println("qidongle")
	file, FileHeader, err := c.Request.FormFile("file")
	//fileType, err := strconv.Atoi(c.PostForm("type")) //获取前端传来的文件的各个参数
	fmt.Println("启动了4")
	if err != nil {
		fmt.Println("获取file err", err)
		return ""
	}
	fmt.Println("启动了5")
	if FileHeader == nil { //这里错了
		fmt.Println("启动了6")

		return ""
	}
	fmt.Println("启动了3")
	FileInfo, err := service.UploadFile(file, FileHeader)
	if err != nil {
		fmt.Println("upload file err:", err)
		return ""
	}
	return "http://127.0.0.1:8081/" + FileInfo.AccessUrl
}
