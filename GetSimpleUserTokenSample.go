package main

import (
	"authenticatin/token"
	"fmt"
)

func main() {
	appId := "1_1"
	appSecret := "77c8d41bd5f44d750d0c02afa2730dc0"
	uid := "uid106"
	expire := 60 * 60 * 24 * 360
	simpleToken, err := token.BuildUserToken(appId, appSecret, int64(expire), uid, token.SIMPLE)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(simpleToken)
}
