package main

import (
	"fmt"
	"github.com/wenlongkang/authenticatin/token"
)

func main() {
	appId := "6_9"
	appSecret := "f04702e9868ff54767fa68ea9ffc8d32"
	uid := "rd107"
	expire := 60 * 60 * 24 * 360
	userToken, err := token.BuildUserToken(appId, appSecret, int64(expire), uid, token.IM)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userToken)
	m, e := token.ParseToken(userToken, appSecret)
	if e != nil {
		fmt.Println("parse token error! ", e)
	}
	appidFromToken := m[token.JWT_TOKEN_APPID].(string)
	uidFromToken := m[token.JWT_TOKEN_UID].(string)
	expFromToken := m[token.JWT_TOKEN_EXP].(float64)
	fmt.Printf("parse token ok.\nappid : %s ,uid %s ,exp %f.\n", appidFromToken, uidFromToken, expFromToken)
}
