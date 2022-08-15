package pkg

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"luntan/global"
	"strings"
)

type FlieType int //filetype =int 用

const TypeImage FlieType = iota + 1 //再次使用时 自动+1

//参数获取方法
func GetFileName(name string) string { //获取文件名称和后缀 处理-；
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	return fileName + ext
}

func GetFileExt(name string) string { //获取后缀
	return path.Ext(name)
}

func GetSavePath() string { //获取文件保存地址
	return global.FileSave.UploadSavePath
}

//检查方法

func CheckSavePath(dst string) bool { //文件保存牡蛎是否存在
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

func CheckContainExit(name string) bool { //检查文件后缀是否包含在约定的配置中
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext) //变大写

	for _, allowExt := range global.FileSave.UploadImageAllowExts {
		if strings.ToUpper(allowExt) == ext {
			fmt.Println("yes")
			return true
		}
	}
	fmt.Println("no")
	return false
}
func CheckMaxSize(f multipart.File) bool { //检查文件大小是否对
	content, _ := ioutil.ReadAll(f)
	size := len(content)

	if size >= global.FileSave.UploadImageMaxSize*1024*1024 {
		return true
	}
	return false
}

func CheckPermission(dst string) bool { //检查文件权限是否足够 雨oserror。errpermis比较
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

//写入方法

func CreateSavePath(dst string, perm os.FileMode) error { //创建保存上传文件的目录
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error { //保存用户上传文件
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer out.Close()
	_, err = io.Copy(out, src) //文件内容拷贝
	return err
}
