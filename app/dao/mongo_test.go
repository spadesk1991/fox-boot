package dao

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/globalsign/mgo/bson"
)

func Test_PipeAll(t *testing.T) {
	t.Run("插入数据", func(t *testing.T) {
		d1 := bson.M{"_id": 1, "name": "jack1", "age": 23, "dept": 1, "score": 89}
		d2 := bson.M{"_id": 2, "name": "jack2", "age": 33, "dept": 1, "score": 46}
		d3 := bson.M{"_id": 3, "name": "jack3", "age": 12, "dept": 2, "score": 78}
		d4 := bson.M{"_id": 4, "name": "jack4", "age": 16, "dept": 2, "score": 90}
		d5 := bson.M{"_id": 5, "name": "jack5", "age": 53, "dept": 3, "score": 70}
		docs := make([]interface{}, 0)
		docs = append(docs, d1, d2, d3, d4, d5)
		if err := Insert("user", docs...); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("聚合查询", func(t *testing.T) {
		pipeline := []bson.M{
			bson.M{
				"$match": bson.M{
					"_id": bson.M{
						"$lte": 4,
					},
				},
			},
			bson.M{
				"$group": bson.M{
					"_id": "$dept",
					"sum_score": bson.M{
						"$sum": "$score",
					},
					"avg_score": bson.M{
						"$avg": "$score",
					},
				},
			},
			bson.M{
				"$skip": 0,
			},
			bson.M{
				"$limit": 2,
			},
		}
		arrRes := make([]bson.M, 0)
		if err := PipeAll("user", pipeline, &arrRes, false); err != nil {
			t.Fatal(err)
		}
		t.Log(arrRes)
	})
	t.Run("删除数据", func(t *testing.T) {
		if err := RemoveAll("user", bson.M{}); err != nil {
			t.Fatal(err)
		}
	})
}

func Test_PipeAll2(t *testing.T) {
	t.Run("插入部门数据", func(t *testing.T) {
		d1 := bson.M{"_id": 1, "dept_name": "IT"}
		d2 := bson.M{"_id": 2, "dept_name": "O&M"}
		docs := make([]interface{}, 0)
		docs = append(docs, d1, d2)
		if err := Insert("dept", docs...); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("插入员工数据", func(t *testing.T) {
		d1 := bson.M{"_id": 1, "name": "jack1", "age": 23, "dept": 1, "score": 89}
		d2 := bson.M{"_id": 2, "name": "jack2", "age": 33, "dept": 1, "score": 46}
		d3 := bson.M{"_id": 3, "name": "jack3", "age": 12, "dept": 2, "score": 78}
		d4 := bson.M{"_id": 4, "name": "jack4", "age": 16, "dept": 2, "score": 90}
		d5 := bson.M{"_id": 5, "name": "jack5", "age": 53, "dept": 3, "score": 70}
		docs := make([]interface{}, 0)
		docs = append(docs, d1, d2, d3, d4, d5)
		if err := Insert("user", docs...); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("连表查询", func(t *testing.T) {
		pipeline := []bson.M{
			bson.M{
				"$lookup": bson.M{
					"from":         "dept",
					"localField":   "dept",
					"foreignField": "_id",
					"as":           "deptArr",
				},
			},
			bson.M{
				"$addFields": bson.M{
					"deptObj": bson.M{
						"$arrayElemAt": []interface{}{"$deptArr", 0},
					},
				},
			},
			bson.M{
				"$project": bson.M{
					"deptArr": 0,
				},
			},
		}
		allRes := make([]bson.M, 0)
		if err := PipeAll("user", pipeline, &allRes, true); err != nil {
			t.Fatal(err)
		}
		bt, _ := json.Marshal(allRes)
		t.Log(string(bt))
	})
	t.Run("清空数据", func(t *testing.T) {
		if err := RemoveAll("user", bson.M{}); err != nil {
			t.Fatal(err)
		}
		if err := RemoveAll("dept", bson.M{}); err != nil {
			t.Fatal(err)
		}
	})

}

func Test_FindAll(t *testing.T) {
	t.Run("查询", func(t *testing.T) {
		arrRes := make([]bson.M, 0)
		if err := FindAll("user", nil, nil, &arrRes); err != nil {
			t.Fatal(err)
		}
		fmt.Println(arrRes)
	})
}
