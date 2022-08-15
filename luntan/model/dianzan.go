package model

import (
	"errors"

	"gorm.io/gorm"
)

type Dianzan struct {
	*gorm.Model
	Postid   int `form:"postid" json:"postid" binding:""`
	Uid      int `form:"uid" json:"uid" binding:""`           //被点赞的帖子的uid
	Followid int `form:"followid" json:"followid" binding:""` //点赞帖子的uid
}

func NewDianzan() Dianzan {
	return Dianzan{}
}

func (g Dianzan) DianzanGet(db *gorm.DB) (Dianzan, error) {
	db = db.Where(
		"id=?",
		g.Uid,
	)
	err := db.Debug().First(&g).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return g, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) { //未关注

	}
	return g, nil
}
func (g Dianzan) DianzanCreate(db *gorm.DB) error {
	return db.Debug().Create(&g).Error
}

//取消关注

func (g Dianzan) DianzanDelete(db *gorm.DB) error {
	return db.Debug().Where("id=?", g.Uid, g.Followid, g.Postid).Delete(&Guanzhu{}).Error
}
