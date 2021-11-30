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
}
