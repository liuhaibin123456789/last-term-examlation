package tool

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"strings"
	"time"
)

var (
	MySecret          = []byte("123456")
	ErrorWrongToken   = errors.New("token格式错误")
	ErrorExpiredToken = errors.New("token过期")
)

const (
	AccessTokenExpireDuration  = time.Hour * 10
	RefreshTokenExpireDuration = time.Hour * 24 * 30
)

type MyClaims struct {
	UserId int64 `json:"user_id,string"`
	jwt.StandardClaims
}

func CreateTwoToken(userId int64) (aToken string, rToken string, err error) {
	//生成accessToken，时间短，保存有用信息
	claim := MyClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenExpireDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "cold bin",
		},
	}
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(MySecret)
	if err != nil {
		return "", "", err
	}
	//生成refreshToken，只需要部门官方字段即可，时间长
	rToken, err = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.StandardClaims{ExpiresAt: time.Now().Add(RefreshTokenExpireDuration).Unix(), Issuer: "cold bin"}).SignedString(MySecret)
	if err != nil {
		return "", "", err
	}
	//将accessToken放到redis里，对应user_id
	err = Set(strconv.FormatInt(userId, 10), aToken, AccessTokenExpireDuration)
	if err != nil {
		return "", "", err
	}
	return aToken, rToken, nil
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		if v.Errors == jwt.ValidationErrorExpired {
			return &MyClaims{}, ErrorExpiredToken
		}
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, ErrorWrongToken
}

// RefreshToken 适用于access_token过期，需要获取新的access_token
func RefreshToken(userId int64, auth string) (newAToken, newRToken string, err error) {
	//拆解auth，拿到refresh_token
	rToken := strings.TrimPrefix(auth, "Bearer ")
	//refresh_token无效直接返回
	if _, err1 := jwt.Parse(rToken, func(t *jwt.Token) (interface{}, error) {
		return MySecret, nil
	}); err1 != nil {
		SugaredError("RefreshToken:", err1)
		err = err1
		return
	}
	return CreateTwoToken(userId)
}
