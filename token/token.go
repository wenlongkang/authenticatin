package token

import (
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const HEADER = "Authorization"
const JwtTokenType = "t"
const JwtTokenUserType = "u"
const JwtTokenAppid = "aid"
const JwtTokenUid = "uid"
const JwtTokenCreatetime = "ct"
const JwtTokenExp = "exp"

func BuildAppToken(appId string, appSecret string, expireSecond int64) (string, error) {
	second := time.Now().Unix() + expireSecond
	cliams := jwt.MapClaims{}
	cliams[JwtTokenType] = APP
	cliams[JwtTokenAppid] = appId
	cliams[JwtTokenExp] = second
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	b, err := base64.StdEncoding.DecodeString(appSecret)
	if err != nil {
		return "", err
	}
	token, err := t.SignedString(b)
	if err != nil {
		return "", err
	}
	return token, err
}

func BuildUserToken(appId string, appSecret string, expireSecond int64, uid string, userType UserType) (string, error) {
	now := time.Now().Unix()
	second := now + expireSecond
	cliams := jwt.MapClaims{}
	cliams[JwtTokenType] = USER
	cliams[JwtTokenAppid] = appId
	cliams[JwtTokenUid] = uid
	cliams[JwtTokenCreatetime] = now * 1000
	cliams[JwtTokenUserType] = userType
	cliams[JwtTokenExp] = second
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	b, err := base64.StdEncoding.DecodeString(appSecret)
	if err != nil {
		return "", err
	}
	token, err := t.SignedString(b)
	if err != nil {
		return "", err
	}
	return token, err
}
