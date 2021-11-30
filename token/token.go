package token

import (
	"encoding/base64"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const HEADER = "Authorization"
const JWT_TOKEN_TYPE = "t"
const JWT_TOKEN_USER_TYPE = "u"
const JWT_TOKEN_APPID = "aid"
const JWT_TOKEN_UID = "uid"
const JWT_TOKEN_CUSTOMERID = "cid"
const JWT_TOKEN_CREATETIME = "ct"
const JWT_TOKEN_EXP = "exp"

func BuildAppToken(appId string, appSecret string, expireSecond int64) (string, error) {
	second := time.Now().Unix() + expireSecond
	cliams := jwt.MapClaims{}
	cliams[JWT_TOKEN_TYPE] = APP
	cliams[JWT_TOKEN_APPID] = appId
	cliams[JWT_TOKEN_EXP] = second
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	b, e := base64.StdEncoding.DecodeString(appSecret)
	if e != nil {
		return "", e
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
	cliams[JWT_TOKEN_TYPE] = USER
	cliams[JWT_TOKEN_APPID] = appId
	cliams[JWT_TOKEN_UID] = uid
	cliams[JWT_TOKEN_CREATETIME] = now * 1000
	cliams[JWT_TOKEN_USER_TYPE] = userType
	cliams[JWT_TOKEN_EXP] = second
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	b, e := base64.StdEncoding.DecodeString(appSecret)
	if e != nil {
		return "", e
	}
	token, err := t.SignedString(b)
	if err != nil {
		return "", err
	}
	return token, err
}

func ParseToken(token string, appSecret string) (map[string]interface{}, error) {
	bytes, e := base64.StdEncoding.DecodeString(appSecret)
	if e != nil {
		return nil, e
	}
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return bytes, nil
	}
	t, e := jwt.Parse(token, keyFunc)
	if e != nil {
		return nil, e
	}
	return t.Claims.(jwt.MapClaims), nil
}
