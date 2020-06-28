package dao

import (
	"github.com/globalsign/mgo"
)

var s *mgo.Session

const mongodb = "test"

func init() {
	var err error
	s, err = mgo.Dial("mongodb://localhost:27017/test")
	s.SetMode(mgo.Monotonic, true)
	if err != nil {
		panic(err)
	}
}

func connect(collection string) (*mgo.Session, *mgo.Collection) {
	ms := s.Copy()
	c := ms.DB(mongodb).C(collection)
	ms.SetMode(mgo.Monotonic, true)
	return ms, c
}

func getDb() (*mgo.Session, *mgo.Database) {
	ms := s.Copy()
	return ms, ms.DB(mongodb)
}

func Count(collection string, query interface{}) (int, error) {
	ms, c := connect(collection)
	defer ms.Close()
	return c.Find(query).Count()
}

func Insert(collection string, docs ...interface{}) error {
	ms, c := connect(collection)
	defer ms.Close()
	return c.Insert(docs...)
}

func FindOne(collection string, query, selector, result interface{}) error {
	ms, c := connect(collection)
	defer ms.Close()

	return c.Find(query).Select(selector).One(result)
}

func FindAll(collection string, query, selector, result interface{}) error {
	ms, c := connect(collection)
	defer ms.Close()

	return c.Find(query).Select(selector).All(result)
}

func FindPage(collection string, page, limit int, query, selector, result interface{}) error {
	ms, c := connect(collection)
	defer ms.Close()

	return c.Find(query).Select(selector).Skip(page * limit).Limit(limit).All(result)
}

func FindIter(collection string, query interface{}) *mgo.Iter {
	ms, c := connect(collection)
	defer ms.Close()

	return c.Find(query).Iter()
}

func Update(collection string, selector, update interface{}) error {
	ms, c := connect(collection)
	defer ms.Close()

	return c.Update(selector, update)
}

func Upsert(collection string, selector, update interface{}) error {
	ms, c := connect(collection)
	defer ms.Close()

	_, err := c.Upsert(selector, update)
	return err
}

func UpdateAll(collection string, selector, update interface{}) error {
	ms, c := connect(collection)
	defer ms.Close()

	_, err := c.UpdateAll(selector, update)
	return err
}

func Remove(collection string, selector interface{}) error {
	ms, c := connect(collection)
	defer ms.Close()

	return c.Remove(selector)
}

func RemoveAll(collection string, selector interface{}) error {
	ms, c := connect(collection)
	defer ms.Close()

	_, err := c.RemoveAll(selector)
	return err
}

//insert one or multi documents
func BulkInsert(collection string, docs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Insert(docs...)
	return bulk.Run()
}

func BulkRemove(collection string, selector ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(collection)
	defer ms.Close()

	bulk := c.Bulk()
	bulk.Remove(selector...)
	return bulk.Run()
}

func BulkRemoveAll(collection string, selector ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.RemoveAll(selector...)
	return bulk.Run()
}

func BulkUpdate(collection string, pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Update(pairs...)
	return bulk.Run()
}

func BulkUpdateAll(collection string, pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.UpdateAll(pairs...)
	return bulk.Run()
}

func BulkUpsert(collection string, pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Upsert(pairs...)
	return bulk.Run()
}

func PipeAll(collection string, pipeline, result interface{}, allowDiskUse bool) error {
	ms, c := connect(collection)
	defer ms.Close()
	var pipe *mgo.Pipe
	if allowDiskUse {
		pipe = c.Pipe(pipeline).AllowDiskUse()
	} else {
		pipe = c.Pipe(pipeline)
	}
	return pipe.All(result)
}

func PipeOne(collection string, pipeline, result interface{}, allowDiskUse bool) error {
	ms, c := connect(collection)
	defer ms.Close()
	var pipe *mgo.Pipe
	if allowDiskUse {
		pipe = c.Pipe(pipeline).AllowDiskUse()
	} else {
		pipe = c.Pipe(pipeline)
	}
	return pipe.One(result)
}
