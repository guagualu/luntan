package global

import (
	"luntan/setting"

	"gorm.io/gorm"
)

var (
	DbEngine      *gorm.DB
	JWTSettings   setting.JWTSetting
	IndexPageSize setting.PinglunindexPager
	XiangQingSize setting.PinglunXiangqingPager
	FileSave      setting.UploadFileSetting
)
