package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"
	"gorm.io/gorm"

	"luntan/global"
	microuser "luntan/microuser/proto"
	"luntan/model"
)

type Microuser struct{}

// Return a new handler
func New() *Microuser {
	return &Microuser{}
}

var dbg *gorm.DB

func Setdbg(db *gorm.DB) {
	dbg = db
}

// Call is a single request handler called via client.Call or the generated client code
func (h *Microuser) Create(ctx context.Context, req *microuser.Createequest, rsp *microuser.CreateResponse) error {
	log.Info("Received usercreate request")
	db := dbg
	log.Info(global.DbEngine)
	u := model.NewUser()
	u.Mail = req.Mail
	u.Name = req.Name
	u.Password = req.Password
	u.PhotoUrl = req.Photourl
	err := db.Debug().Create(&u).Error
	if err != nil {
		rsp.Flag = -1
		return err
	}
	rsp.Id = int64(u.ID)
	rsp.Flag = 1
	return nil
}

func (h *Microuser) Get(ctx context.Context, req *microuser.Usergetrequest, rsp *microuser.Usergetresponse) error {
	log.Info("Received userget request")
	db := dbg
	db = db.Where(
		"id=?",
		req.Id,
	)
	u := model.NewUser()
	err := db.Debug().First(&u).Error
	if err != nil {
		return err
	}
	rsp.Id = int64(u.ID)
	rsp.Password = u.Password
	rsp.Mail = u.Mail
	rsp.Photourl = u.PhotoUrl
	rsp.Name = u.Name
	rsp.Dianzanshu = int64(u.Dianzanshu)
	return nil
}
func (h *Microuser) Dengluget(ctx context.Context, req *microuser.Userdenglugetrequest, rsp *microuser.Userdenglugetresponse) error {
	log.Info("Received userdengluget request")
	db := dbg
	db = db.Where(
		"id=? and password=?",
		req.Id, req.Password,
	)
	u := model.NewUser()
	err := db.Debug().First(&u).Error
	if err != nil {
		return err
	}
	rsp.Id = int64(u.ID)
	rsp.Password = u.Password
	rsp.Mail = u.Mail
	rsp.Photourl = u.PhotoUrl
	rsp.Name = u.Name
	rsp.Dianzanshu = int64(u.Dianzanshu)
	return nil
}

func (h *Microuser) Update(ctx context.Context, req *microuser.Updaterequest, rsp *microuser.Updateresponse) error {
	log.Info("Received userupdate request")
	db := dbg
	u := model.NewUser()
	// u.ID=uint(req.Id)
	u.Name = req.Name
	u.Mail = req.Mail
	u.Password = req.Password
	u.PhotoUrl = req.Photourl
	err := db.Debug().Model(&model.User{}).Where("id=?", req.Id).Updates(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (h *Microuser) Paihang(ctx context.Context, req *microuser.Userpaihangrequest, rsp *microuser.Userpaihangresponse) error {
	log.Info("Received userpaihang request")
	db := dbg
	// u.ID=uint(req.Id)
	users := []*model.User{}
	err := db.Debug().Order("dianzanshu desc").Limit(10).Offset((int(req.Page) - 1) * 10).Find(&users).Error
	if err != nil {
		return err
	}
	for i := 0; i < len(users); i++ {
		rsp.List = append(rsp.List, &microuser.Userpaihangresponse_List{Name: users[i].Name, Id: int64(users[i].ID), Mail: users[i].Mail, Photourl: users[i].PhotoUrl, Dianzanshu: int64(users[i].Dianzanshu)})
	}
	return nil
}
