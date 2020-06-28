package service

import (
	"LiteService/app/dao"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
)

type TestService struct{}

const (
	TestCollection = "test_collections"
)

// Create 插入数据
func (t *TestService) Create(docs map[string]interface{}) error {
	return dao.Exec(TestCollection, func(c *mgo.Collection) (e error) {
		return c.Insert(docs)
	})
}

func (t *TestService) list(query bson.M, res []map[string]interface{}) error {
	return dao.Exec(TestCollection, func(c *mgo.Collection) (e error) {
		return c.Find(query).All(&res)
	})
}

func (t *TestService) Agge(query bson.M, res interface{}) error {
	return dao.Exec(TestCollection, func(c *mgo.Collection) (e error) {
		//c.Find(query).
		return
	})
}
