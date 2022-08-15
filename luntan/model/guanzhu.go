package model

import (
	"errors"

	"gorm.io/gorm"
)

type Guanzhu struct {
	*gorm.Model
	Uid      int `form:"uid" json:"uid" binding:""` //被关注
	Followid int `form:"followid" json:"followid" binding:""`
}

func NewGuanzhu() Guanzhu {
	return Guanzhu{}
}

func (g Guanzhu) GuanzhuGet(db *gorm.DB) (Guanzhu, error) {
	db = db.Where(
		"uid=? and followid=?",
		g.Uid, g.Followid,
	)
	err := db.Debug().First(&g).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return g, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) { //未关注

	}
	return g, nil
}
func (g Guanzhu) GuanzhuList(db *gorm.DB) ([5]*Guanzhu, error) {
	var g1 [5]*Guanzhu
	db = db.Where(
		"uid=?",
		g.Uid,
	)
	err := db.Debug().Limit(5).Find(&g1).Error
	if err != nil {
		return g1, err
	}
	return g1, nil
}
func (g Guanzhu) GuanzhuCreate(db *gorm.DB) error {
	return db.Debug().Create(&g).Error
}

//取消关注

func (g Guanzhu) GuanzhuDelete(db *gorm.DB) error {
	return db.Debug().Where("uid=? and followid=?", g.Uid, g.Followid).Delete(&Guanzhu{}).Error
}
