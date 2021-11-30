package main

import (
	"fmt"
	"github.com/wenlongkang/authenticatin/token"
)

func main() {
	appId := "11_55"
	appSecret := "ff7dc971c261636f7cc3e774b7cd7d85"
	uid := "1157uid101"
	expire := 60 * 60 * 24 * 360
	simpleToken, err := token.BuildUserToken(appId, appSecret, int64(expire), uid, token.TRUST)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(simpleToken)
}
