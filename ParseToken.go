package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
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
	claims, e := token.ParseToken(userToken, appSecret)
	if e != nil {
		fmt.Println("parse token error! ", e)
	}
	appidFromToken := claims.(jwt.MapClaims)[token.JWT_TOKEN_APPID]
	uidFromToken := claims.(jwt.MapClaims)[token.JWT_TOKEN_UID]
	expFromToken := claims.(jwt.MapClaims)[token.JWT_TOKEN_EXP]
	fmt.Printf("parse token ok.\nappid : %s ,uid %s ,exp %f.\n", appidFromToken, uidFromToken, expFromToken)
}
