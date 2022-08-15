package handler

import (
	"context"
	"errors"

	log "github.com/micro/micro/v3/service/logger"
	"gorm.io/gorm"

	microdianzan "luntan/microdianzan/proto"
	"luntan/model"
)

type Microdianzan struct{}

// Return a new handler
func New() *Microdianzan {
	return &Microdianzan{}
}

var dbg *gorm.DB

func Setdbg(db *gorm.DB) {
	dbg = db
}

// Call is a single request handler called via client.Call or the generated client code
func (h *Microdianzan) Create(ctx context.Context, req *microdianzan.Createequest, rsp *microdianzan.CreateResponse) error {
	log.Info("Received dianzancreate request")
	db := dbg
	g := model.NewDianzan()
	g.Uid = int(req.Uid)
	g.Postid = int(req.Postid)
	g.Followid = int(req.Followid)
	err := db.Debug().Create(&g).Error
	if err != nil {
		return err
	}
	return nil
}

func (h *Microdianzan) Get(ctx context.Context, req *microdianzan.Dianzangetrequest, rsp *microdianzan.Dianzangetresponse) error {
	log.Info("Received dianzanget request")
	db := dbg
	db = db.Where(
		"uid=? and followid=? and postid=?",
		req.Uid, req.Followid, req.Postid,
	)
	g := model.NewDianzan()
	err := db.Debug().First(&g).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		rsp.Flag = -1
		return nil
	}
	if err != nil {
		return err
	}
	rsp.Id = int64(g.ID)
	rsp.Uid = int64(g.Uid)
	rsp.Followid = int64(g.Followid)
	rsp.Postid = int64(g.Postid)
	rsp.Flag = 1
	return nil
}

func (h *Microdianzan) Delete(ctx context.Context, req *microdianzan.Dianzandeletrequest, rsp *microdianzan.Dianzandeleteresponse) error {
	log.Info("Received dianzandelete request")
	db := dbg
	err := db.Debug().Where("id=?", req.Id).Delete(&model.Dianzan{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (h *Microdianzan) Dianzan(ctx context.Context, req *microdianzan.DianzanReq, rsp *microdianzan.DianzanRsp) error {
	log.Info("Received dianzan request")
	log.Info("req", req)
	db := dbg
	err := db.Transaction(func(tx *gorm.DB) error {
		var flag int
		tx1 := tx.Where(
			"uid=? and postid=?",
			req.Uid, req.Postid,
		)
		g := model.NewDianzan()
		err := tx1.Debug().First(&g).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			flag = -1
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if flag == -1 {
			g1 := model.NewDianzan()
			g1.Uid = int(req.Uid)
			g1.Postid = int(req.Postid)
			g1.Followid = int(req.Followid)
			tx2 := tx
			err := tx2.Debug().Create(&g1).Error
			if err != nil {
				return err
			}
			t1 := model.NewTiezi()
			txpostget := tx
			err = txpostget.Debug().Where("ID =?", req.Postid).First(&t1).Error
			if err != nil && err != gorm.ErrRecordNotFound {
				return err
			}
			dianzanshu := t1.Dianzanshu + 1
			tx3 := tx
			err = tx3.Model(&t1).Debug().Where("ID =?", req.Postid).Update("dianzanshu", dianzanshu).Error
			if err != nil {
				return err
			}
			u1 := model.NewUser()
			txuserget := tx
			err = txuserget.Debug().Where("ID =?", req.Uid).First(&u1).Error
			if err != nil && err != gorm.ErrRecordNotFound {
				return err
			}
			dianzanshu1 := u1.Dianzanshu + 1
			tx4 := tx
			err = tx4.Model(&u1).Debug().Update("dianzanshu", dianzanshu1).Error
			if err != nil {
				return err
			}
		} else {
			g1 := model.NewDianzan()
			g1.Uid = int(req.Uid)
			g1.Postid = int(req.Postid)
			g1.Followid = int(req.Followid)
			tx2 := tx
			err := tx2.Debug().Delete(&g).Error
			if err != nil {
				return err
			}
			t1 := model.NewTiezi()
			txpostget := tx
			err = txpostget.Debug().Where("ID =?", req.Postid).First(&t1).Error
			if err != nil && err != gorm.ErrRecordNotFound {
				return err
			}
			dianzanshu := t1.Dianzanshu - 1
			tx3 := tx
			err = tx3.Model(&t1).Debug().Where("ID =?", req.Postid).Update("dianzanshu", dianzanshu).Error
			if err != nil {
				return err
			}
			u1 := model.NewUser()
			txuserget := tx
			err = txuserget.Debug().Where("ID =?", req.Uid).First(&u1).Error
			if err != nil && err != gorm.ErrRecordNotFound {
				return err
			}
			dianzanshu1 := u1.Dianzanshu - 1
			tx4 := tx
			err = tx4.Model(&u1).Debug().Update("dianzanshu", dianzanshu1).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	rsp.Flag = 1
	return nil
}
