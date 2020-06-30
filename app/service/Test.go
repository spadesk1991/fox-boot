package service

import (
	"LiteService/app/dao"

	"github.com/globalsign/mgo/bson"
)

type TestService struct{}

const (
	TestCollection = "test_collections"
)

// Create 插入数据
func (t *TestService) Create(docs map[string]interface{}) error {
	return dao.Insert(TestCollection, docs)
}

func (t *TestService) list(query bson.M, res []map[string]interface{}) error {
	return dao.FindAll(TestCollection, query, nil, &res)
}

func (t *TestService) Agge(query bson.M, res interface{}) error {
	return nil
}
