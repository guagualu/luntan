package service

import (
	"errors"
	"fmt"
	"luntan/global"
	"luntan/pkg"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

//上传文件小封装
func UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fmt.Println("启动了2")
	filename := pkg.GetFileName(fileHeader.Filename)
	uploadSavePath := pkg.GetSavePath()
	dst := uploadSavePath + "/" + filename
	if !pkg.CheckContainExit(filename) {
		return nil, errors.New("file suffix is not support")
	}
	if !pkg.CheckSavePath(uploadSavePath) {
		err := pkg.CreateSavePath(uploadSavePath, os.ModePerm)
		if err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	if pkg.CheckMaxSize(file) {
		return nil, errors.New("exceeded maximum file limit")
	}
	if pkg.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions")
	}
	if err := pkg.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}
	AccessUrl := global.FileSave.UploadSavePath + "/" + filename //*******************因为没有上线 所以地址直接从本地取
	return &FileInfo{Name: filename, AccessUrl: AccessUrl}, nil
}
