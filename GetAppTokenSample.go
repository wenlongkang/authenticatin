package main

import (
	"fmt"
	"github.com/wenlongkang/authenticatin/token"
)

func main() {
	appId := "5_87"
	appSecret := "b5a818c62ef89b4d805b308d85b6732d"
	expire := 60 * 60 * 24 * 360 // 360 天
	appToken, err := token.BuildAppToken(appId, appSecret, int64(expire))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(appToken)
}
