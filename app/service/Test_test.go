package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/globalsign/mgo/bson"
)

func TestTestService_Agge(t1 *testing.T) {
	type args struct {
		query bson.M
		res   interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TestService{}
			if err := t.Agge(tt.args.query, tt.args.res); (err != nil) != tt.wantErr {
				t1.Errorf("Agge() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTestService_Create(t1 *testing.T) {
	type args struct {
		docs map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "新增数据1", args: args{docs: map[string]interface{}{
			"name":       "jack1",
			"age":        23,
			"created_at": time.Now(),
		}}, wantErr: false},
		{name: "新增数据2", args: args{docs: map[string]interface{}{
			"name":       "jack2",
			"age":        13,
			"created_at": time.Now(),
		}}, wantErr: false},
		{name: "新增数据3", args: args{docs: map[string]interface{}{
			"name":       "jack3",
			"age":        3,
			"created_at": time.Now(),
		}}, wantErr: false},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TestService{}
			if err := t.Create(tt.args.docs); (err != nil) != tt.wantErr {
				t1.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTestService_list(t1 *testing.T) {
	type args struct {
		query bson.M
		res   []map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "根据时间查询", args: args{
			query: bson.M{
				"created_at": bson.M{
					"$lt": time.Now(),
				},
			},
			res: nil,
		}, wantErr: false},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TestService{}
			if err := t.list(tt.args.query, tt.args.res); (err != nil) != tt.wantErr {
				fmt.Println(tt.args.res)
				t1.Errorf("list() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
