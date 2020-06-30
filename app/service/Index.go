package service

import (
	"LiteService/app/common"
	"fmt"
	"os"
)

type regService struct {
	url  string
	apis map[string][]string
}

func NewRegService(prefix string) *regService {
	return &regService{
		url: "http://172.17.0.1:88/sign",
		apis: map[string][]string{
			fmt.Sprintf("http://172.17.0.1:%s", os.Getenv("port")): []string{fmt.Sprintf("%s/*", prefix)},
		},
	}
}

func (r regService) Reg() {
	if os.Getenv("GIN_MODE") != "" {
		res, err := common.DefaultRQ().Uri(r.url).Post().SetBody(r.apis).StringResult()
		if err != nil {
			panic(err)
		}
		fmt.Printf("registry router ok, at %v\n", res)
	}
}
