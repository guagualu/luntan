package model

import (
	"gorm.io/gorm"
)

//暂时在这些 等后面每个为服务做好后移过去
type User struct {
	gorm.Model
	Name       string `json:"name" form:"name" binding:""`
	Password   string `json:"password" form:"password" binding:""`
	Mail       string `json:"mail" form:"mail" binding:""`
	Dianzanshu int    `json:"dianzanshu" form:"dianzanshu" binding:"" gorm:"index:dianzanindex"` //还做个索引最好
	PhotoUrl   string `json:"photourl" form:"photourl" binding:""`
}

func NewUser() User {
	return User{}
}

func (u User) UserGet(db *gorm.DB) (User, error) {
	db = db.Where(
		"id=?",
		u.ID,
	)
	err := db.Debug().First(&u).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u, err
	}
	return u, nil
}
func (u User) UserCreate(db *gorm.DB) error {
	return db.Debug().Create(&u).Error
}

func (u User) UserUpdate(db *gorm.DB) error {
	return db.Debug().Model(&User{}).Where("id=?", u.ID).Updates(&u).Error
}

func (u User) UserPaihang(db *gorm.DB) ([10]*User, error) {
	users := [10]*User{}
	err := db.Debug().Order("dianzanshu desc").Find(&users).Limit(10).Error
	if err != nil {
		return users, err
	}
	return users, nil
}
