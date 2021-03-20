package app

import (
	"chttp/global"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims jwt结构体
type Claims struct {
	AppKey    string `JSON:"app_key"`
	AppSecret string `JSON:"app_secret"`
	UID       int    `JSON:"uid"`
	jwt.StandardClaims
}

// 获取签名
func getJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

// GenerateToken 创建 token
func GenerateToken(uid int) (string, error) {

	nowTime := time.Now()
	// m, _ := time.ParseDuration("24h")
	expiceTime := nowTime.Add(global.JWTSetting.Expire)
	// claims := Claims{
	// 	AppKey:    util.EncodeMD5("陈大爷"),
	// 	AppSecret: util.EncodeMD5("chendaye"),
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: expiceTime.Unix(),
	// 		Issuer:    "陈大爷",
	// 		Id:        strconv.Itoa(int(id)),
	// 	},
	// }
	claims := jwt.MapClaims{
		// "AppKey":    util.EncodeMD5("陈大爷"),
		// "AppSecret": util.EncodeMD5("chendaye"),
		"uid": uid,
		"exp": expiceTime.Unix(),
		"iss": global.JWTSetting.Issuer,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(getJWTSecret()) // 这里就是
	return token, err
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return getJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
