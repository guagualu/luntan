package pkg

import (
	"fmt"
	"luntan/global"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Id                 string `json:"id"`
	Secret             string `json:"secret"`
	jwt.StandardClaims        //JWT自定义预定库
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSettings.Secret)
}

func GenerateToken(id, secret string) (string, error) { //生成token
	nowtime := time.Now()
	expireTime := nowtime.Add(global.JWTSettings.Expire)
	claims := Claims{
		// AppKey:    util.EncodeMD5(appKey),
		// AppSecret: util.EncodeMD5(appSecret), //加密 jwt自己会加密 先去掉试试
		Id:     id,
		Secret: secret,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSettings.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //后面来看  根据claims结构体创建token  有加密算法
	token, err := tokenClaims.SignedString(GetJWTSecret())           //根据密匙生成token
	//token, err := tokenClaims.SignedString([]byte("eddycjy"))
	// fmt.Println(err)
	return token, err

}

func ParseToken(token string) (*Claims, error) { //验证token
	claim := &Claims{}
	tokenClaims, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) { //解码校验
		return GetJWTSecret(), nil
		//return []byte("thegua"), nil

	})
	if tokenClaims != nil {
		claims, _ := tokenClaims.Claims.(*Claims)
		fmt.Printf("解析出来的%v\n", claims)
		// if ok && tokenClaims.Valid { //若过期时间等验证
		// 	fmt.Println("getclaims成功")
		// 	return claims, nil
		// }
		return claims, nil
	}
	return nil, err
}
