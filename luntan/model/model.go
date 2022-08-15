package model

import (
	"fmt"
	"luntan/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBModel(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:13306)/%s?charset=%s&parseTime=%t&loc=Local", databaseSetting.UserName, databaseSetting.Password, databaseSetting.DBName, databaseSetting.Charset, databaseSetting.ParseTime)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&TieziPinglun{})
	//db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(172.17.0.2:13306)/%s?charset=%s&parseTime=%t&loc=Local", databaseSetting.UserName, databaseSetting.Password, databaseSetting.DBName, databaseSetting.Charset, databaseSetting.ParseTime)) //有关数据库的信息来连接
	if err != nil {
		panic(err)
	}
	return db, nil
}
