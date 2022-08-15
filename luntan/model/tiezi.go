package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Tiezi struct {
	// PostcommentID string        `gorm:"primary key" json:"postcommentid" form:"postcommentid" binding:"required"`
	// PostId        string        `json:"postid" form:"postid" binding:"required"`
	// Time          time.Duration `json:"time" form:"time"`
	// Id            string        `json:"id" form:"id"`
	// Content       string        `json:"content" form:"content"`
	*gorm.Model
	Userid     int    `json:"userid" form:"userid"`
	Content    string `json:"content" form:"content"`
	Title      string `json:"title" form:"title"`
	Photourl   string `json:"photourl" form:"photourl"`
	Dianzanshu int    `json:"dianzanshu" form:"dianzanshu"`
}
type TieziPinglun struct {
	*gorm.Model
	Postid  int    `json:"postid" form:"postid" binding:"required"`
	Time    string `json:"time" form:"time"`
	Uid     int    `json:"uid" form:"uid"`
	Content string `json:"content" form:"content"`
}

func NewTiezi() Tiezi {
	return Tiezi{}
}
func NewTieziPinglun() TieziPinglun {
	return TieziPinglun{}
}
func (t Tiezi) TieziGet(db *gorm.DB) (Tiezi, error) {
	db = db.Where(
		"id=?",
		t.ID,
	)
	err := db.Debug().First(&t).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("帖子查询错误")
		return t, err
	}
	return t, nil
}
func TieziList(pageOffset, pageSize int, db *gorm.DB) ([]*Tiezi, error) {
	var tags []*Tiezi
	var err error
	db = db.Offset(pageOffset).Limit(pageSize) //==============后面前
	// fmt.Println("gua")
	// fmt.Println(t.Tablename()) //自己加的 测试是否自动运行了 知道是哪个表
	db = db.Order("created_at desc")
	if err = db.Debug().Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil

}
func (t Tiezi) TieziCreate(db *gorm.DB) error {
	return db.Create(&t).Error
}
func (t Tiezi) TieziUpdate(db *gorm.DB) error {
	db = db.Where("id=?", t.ID)                                                            //限制条件 db具有流式 model 指定db操作的实力模型 一般在没有find 等指定类型时使用 操作限制条件是传进来的tag的数据行
	err := db.Debug().Model(&Tiezi{}).Select("content").Update("content", t.Content).Error //用save不报错 create报错 只能先将就save 用了 配合selct 也解决了 0值不能改的问题
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (t Tiezi) TieziDelete(db *gorm.DB) error {
	return db.Debug().Where("tiezibianhao", t.ID).Delete(&t).Error
}

//**************评论部分
func (t TieziPinglun) TieziPinglunGet(db *gorm.DB) (TieziPinglun, error) {
	db = db.Where(
		"id=?",
		t.ID,
	)
	err := db.First(&t).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("帖子查询错误")
		return t, err
	}
	return t, nil
}
func (t TieziPinglun) TieziPinglunList(pageOffset, pageSize int, db *gorm.DB) ([]TieziPinglun, error) {
	var tags []TieziPinglun
	var err error
	if pageOffset >= 0 && pageOffset > 0 {
		db = db.Debug().Offset(pageOffset).Limit(pageSize) //==============后面前
	}
	// fmt.Println("gua")
	// fmt.Println(t.Tablename()) //自己加的 测试是否自动运行了 知道是哪个表
	// if t. != "" {
	// 	db.Where("name=?", t.Title)
	// }
	// db = db.Where("state=?", t.)
	db = db.Order("created_at desc").Where("id=?", t.Postid)
	if err = db.Debug().Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil

}
func (t TieziPinglun) TieziPinglunCreate(db *gorm.DB) error {
	return db.Create(&t).Error
}
func (t TieziPinglun) TieziPinglunUpdate(db *gorm.DB) error {
	db = db.Model(&TieziPinglun{}).Where("id=?", t.ID)                     //限制条件 db具有流式 model 指定db操作的实力模型 一般在没有find 等指定类型时使用 操作限制条件是传进来的tag的数据行
	err := db.Debug().Select("comment").Update("content", t.Content).Error //用save不报错 create报错 只能先将就save 用了 配合selct 也解决了 0值不能改的问题
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (t TieziPinglun) TieziPinglunDelete(db *gorm.DB) error {
	return db.Debug().Where("id=?", t.ID).Delete(&t).Error
}
