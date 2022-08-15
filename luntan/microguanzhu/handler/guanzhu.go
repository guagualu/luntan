package handler

import (
	"context"
	"errors"

	log "github.com/micro/micro/v3/service/logger"
	"gorm.io/gorm"

	microguanzhu "luntan/microguanzhu/proto"
	helloworld "luntan/microhandler/proto"
	"luntan/model"
)

type Microguanzhu struct{}

// Return a new handler
func New() *Microguanzhu {
	return &Microguanzhu{}
}

var dbg *gorm.DB

func Setdbg(db *gorm.DB) {
	dbg = db
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Microguanzhu) Call(ctx context.Context, req *helloworld.Request, rsp *helloworld.Response) error {
	log.Info("Received Helloworld.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
func (h *Microguanzhu) Create(ctx context.Context, req *microguanzhu.Createequest, rsp *microguanzhu.CreateResponse) error {
	log.Info("Received guanzhucreate request")
	db := dbg
	g := model.NewGuanzhu()
	g.Uid = int(req.Uid)
	g.Followid = int(req.Followid)
	err := db.Debug().Create(&g).Error
	if err != nil {
		rsp.Flag = -1
		return err
	}
	rsp.Flag = 1
	return nil
}

func (h *Microguanzhu) Get(ctx context.Context, req *microguanzhu.Guanzhugetrequest, rsp *microguanzhu.Guanzhugetresponse) error {
	log.Info("Received guanzhuget request")
	db := dbg
	db = db.Where(
		"uid=? and followid=?",
		req.Uid, req.Followid,
	)
	g := model.NewGuanzhu()
	err := db.Debug().First(&g).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) { //未关注
		rsp.Flag = -1
		return nil
	}
	rsp.Id = int64(g.ID)
	rsp.Uid = int64(g.Uid)
	rsp.Followid = int64(g.Followid)
	rsp.Flag = 1
	return nil
}

func (h *Microguanzhu) List(ctx context.Context, req *microguanzhu.Guanzhulistrequest, rsp *microguanzhu.Guanzhulistresponse) error {
	log.Info("Received guanzhulist request")
	db := dbg
	var g1 []*model.Guanzhu
	db = db.Where(
		"uid=?",
		req.Uid,
	)
	err := db.Debug().Limit(10).Offset((int(req.Page) - 1) * 10).Find(&g1).Error
	if err != nil {
		return err
	}
	for i := 0; i < len(g1); i++ {
		rsp.List = append(rsp.List, &microguanzhu.Guanzhulistresponse_List{Uid: int64(g1[i].Uid), Followid: int64(g1[i].Followid), Id: int64(g1[i].ID)})
	}
	return nil
}

func (h *Microguanzhu) Delete(ctx context.Context, req *microguanzhu.Guanzhudeletrequest, rsp *microguanzhu.Guanzhudeleteresponse) error {
	log.Info("Received guanzhudelete request")
	db := dbg
	err := db.Debug().Where("uid=? and followid=?", req.Uid, req.Followid).Delete(&model.Guanzhu{}).Error
	if err != nil {
		return err
	}
	return nil
}
