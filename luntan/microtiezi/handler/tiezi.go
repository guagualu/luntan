package handler

import (
	"context"
	"fmt"

	log "github.com/micro/micro/v3/service/logger"
	"gorm.io/gorm"

	microtiezi "luntan/microtiezi/proto"
	"luntan/model"
)

type Microtiezi struct{}

// Return a new handler
func New() *Microtiezi {
	return &Microtiezi{}
}

var dbg *gorm.DB

func Setdbg(db *gorm.DB) {
	dbg = db
}

// Call is a single request handler called via client.Call or t
func (h *Microtiezi) Create(ctx context.Context, req *microtiezi.Createtieziquest, rsp *microtiezi.CreatetieziResponse) error {
	log.Info("Received tiezicreate request")
	db := dbg
	t := model.NewTiezi()
	t.Userid = int(req.Userid)
	t.Content = req.Content
	t.Photourl = req.Photourl
	t.Title = req.Title
	t.Dianzanshu = int(req.Dianzanshu)
	err := db.Debug().Create(&t).Error
	if err != nil {
		rsp.Flag = -1
		return err
	}
	rsp.Id = int64(t.ID) //可以获取自增的id
	rsp.Flag = 1
	return nil
}

func (h *Microtiezi) Get(ctx context.Context, req *microtiezi.Tiezigetrequest, rsp *microtiezi.Tiezigetresponse) error {
	log.Info("Received tieziget request")
	db := dbg
	db = db.Where(
		"id=?",
		req.Id,
	)
	t := model.NewTiezi()
	err := db.Debug().First(&t).Error
	if err != nil {
		return err
	}
	rsp.Id = int64(t.ID)
	rsp.Title = t.Title
	rsp.Content = t.Content
	rsp.Photourl = t.Photourl
	rsp.Userid = int64(t.Userid)
	rsp.Dianzanshu = int64(t.Dianzanshu)
	return nil
}

func (h *Microtiezi) List(ctx context.Context, req *microtiezi.Tiezilistrequest, rsp *microtiezi.TiezilistResponse) error {
	log.Info("Received tiezilist request")
	db := dbg
	var tags []*model.Tiezi
	var err error
	db = db.Offset((int(req.Pageoffset) - 1) * 10).Limit(10) //==============后面前
	// fmt.Println("gua")
	// fmt.Println(t.Tablename()) //自己加的 测试是否自动运行了 知道是哪个表
	db = db.Order("created_at desc")
	if err = db.Debug().Find(&tags).Error; err != nil {
		return err
	}
	fmt.Println(tags)
	for i := 0; i < len(tags); i++ {
		m := &microtiezi.TiezilistResponse_List{Id: int64(tags[i].ID), Title: tags[i].Title, Content: tags[i].Content, Photourl: tags[i].Photourl, Dianzanshu: int64(tags[i].Dianzanshu), Userid: int64(tags[i].Userid)}
		rsp.List = append(rsp.List, m)
	}

	return nil

}

func (h *Microtiezi) Createpinglun(ctx context.Context, req *microtiezi.Createtiezipinglunquest, rsp *microtiezi.CreatetiezipinglunResponse) error {
	log.Info("Received tiezipingluncreate request")
	db := dbg
	t := model.NewTieziPinglun()
	t.Content = req.Content
	t.Uid = int(req.Uid)
	t.Postid = int(req.Postid)
	err := db.Debug().Create(&t).Error
	if err != nil {
		rsp.Flag = -1
		return err
	}
	rsp.Flag = 1
	return nil
}

func (h *Microtiezi) Listpinglun(ctx context.Context, req *microtiezi.Tiezipinglunlistrequest, rsp *microtiezi.TiezipinglunlistResponse) error {
	log.Info("Received tiezipinglunlist request")
	db := dbg
	var tags []*model.TieziPinglun
	var err error
	db = db.Order("created_at desc").Where("postid=?", req.Postid)
	if err = db.Debug().Limit(10).Offset((int(req.Page) - 1) * 10).Find(&tags).Error; err != nil {
		return err
	}
	for i := 0; i < len(tags); i++ {
		m := &microtiezi.TiezipinglunlistResponse_List{Postid: int64(tags[i].Postid), Uid: int64(tags[i].Uid), Content: tags[i].Content}
		rsp.List = append(rsp.List, m)
	}

	return nil
}
