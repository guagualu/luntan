package main

import (
	"log"
	"luntan/global"
	"luntan/model"
	"luntan/router"
	"luntan/setting"
)

//**********************88注意看有这种样式的注释
// docker run -it --network host --rm mysql mysql -h127.0.0.1 -P13306 --default-character-set=utf8mb4 -uroot -p
// 172.17.0.2:3306
// docker run -it --network host --rm mysql mysql -h172.17.0.2 -P3306 --default-character-set=utf8mb4 -uroot -p
// gin.SetMode(global.ServerSetting.RunMode) //属于server的配置模板
// 	r := routers.NewRouter()
// 	s := &http.Server{
// 		Addr:              ":" + global.ServerSetting.HttpPort,
// 		Handler:           r, //注--------
// 		ReadHeaderTimeout: global.ServerSetting.ReadTimeout,
// 		WriteTimeout:      global.ServerSetting.WriteTimeout,
// 		MaxHeaderBytes:    1 << 20,
// 	} //大写外用
// 	global.Logger.Infof("%s: do-programming-tour-book/ %s", "eddycjy", "blog_service") //日志test
// 	s.ListenAndServe()
//***********************文件需要file type  分页需要page

//点赞要用个mq  还未完成
//点赞排行
//文件
//日志用原生
//所有服务都用rpc方式
//redis后面再看
//关注的关联表 :一张单独的关联表表  id 与 followid

// 还有注册返回id 后面改
// 登陆还需要前端存cookie token  还要再开一个全局 js 传token
//字段必须带json 否则传到前端无法解析
// node server.js

//gorm 自增的 create 的 会反馈到绑定的变量上
func main() {
	dbs := setting.SetDB()
	global.JWTSettings = setting.SetJWT()
	global.FileSave = setting.SetUploadFile()
	global.IndexPageSize = setting.SetIPage()
	Dbengine, err := model.NewDBModel(dbs)
	if err != nil {
		log.Println(err)
	}
	global.DbEngine = Dbengine
	r := router.NewRouter()
	r.Run()
}
