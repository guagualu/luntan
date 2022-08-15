package setting

import "time"

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleconns int
	MaxOpenConns int
}
type UploadFileSetting struct {
	UploadSavePath       string
	UploadServerUrl      string
	UploadImageMaxSize   int
	UploadImageAllowExts [2]string
}
type JWTSetting struct {
	Secret string
	Issuer string //颁发者唯一标识
	Expire time.Duration
}
type PinglunindexPager struct {
	DefaultPageSize int64 //每页固定多少条 没有的就为空
}
type PinglunXiangqingPager struct {
	DefaultPageSize int64 //每页固定多少条 没有的就为空
}

func SetDB() *DatabaseSettingS {
	db := DatabaseSettingS{DBType: "mysql", UserName: "root", Password: "root1234", Host: "127.0.0.1:13306", DBName: "luntan", Charset: "utf8mb4", ParseTime: true}
	// db := DatabaseSettingS{DBType: "mysql", UserName: "root", Password: "root1234", Host: "172.17.0.2:3306", DBName: "shenqingsystem", Charset: "utf8mb4", ParseTime: true}
	//db := DatabaseSettingS{DBType: "mysql", UserName: "gymApplication", Password: "12345678", Host: "120.78.170.43:3306", DBName: "gymapplication", Charset: "utf8mb4", ParseTime: true}
	return &db
}
func SetJWT() JWTSetting {
	jwt := JWTSetting{Secret: "thegua", Issuer: "thegua", Expire: 7200}
	return jwt
}
func SetXPage() PinglunXiangqingPager {
	p := PinglunXiangqingPager{DefaultPageSize: 2}
	return p
}
func SetIPage() PinglunindexPager {
	p := PinglunindexPager{DefaultPageSize: 2}
	return p
}
func SetUploadFile() UploadFileSetting {
	p := UploadFileSetting{UploadSavePath: "storage/uploads", UploadServerUrl: "http://127.0.0.1:8000/static", UploadImageMaxSize: 5}
	p.UploadImageAllowExts[0] = ".jpg"
	p.UploadImageAllowExts[1] = ".png"
	return p
}
