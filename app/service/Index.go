package service

import (
	"LiteService/app/common"
	"fmt"
	"os"
)

func Reg() {
	if os.Getenv("GIN_MODE") != "" {
		url := "http://172.17.0.1:88/sign"
		apis := map[string][]string{
			fmt.Sprintf("http://172.17.0.1:%s", os.Getenv("port")): []string{"/api/demo"},
		}
		res, err := common.DefaultRQ().Uri(url).Post().SetBody(apis).StringResult()
		if err != nil {
			panic(err)
		}
		fmt.Printf("registry router ok, at %v\n", res)
	}
}
