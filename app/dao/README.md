   ## mongodb
   ## mongodb 是什么
    MongoDB：是一个数据库 ,高性能、无模式、文档性，目前nosql中最热门的数据库，开源产品，基于c++开发。是nosql数据库中功能最丰富，最像关系数据库的
   ### 特性
  1. 面向集合文档的存储：适合存储Bson（json的扩展）形式的数据；
  2. 格式自由，数据格式不固定，生产环境下修改结构都可以不影响程序运行；
  3. 强大的查询语句，面向对象的查询语言，基本覆盖sql语言所有能力；
  4. 完整的索引支持，支持查询计划；
  5. 支持复制和自动故障转移；
  6. 支持二进制数据及大型对象（文件）的高效存储；
  7. 使用内存映射存储引擎，把磁盘的IO操作转换成为内存的操作；
  ### mongodb 与 mysql 对应关系
 | mongodb|mysql|
 |:-----|:---|
 |db|database|
 |collection|table|
 |document|row/record|
 |field|column|
 |$lookup 或 嵌入文档|table joins|
 |primary key|_id|
 ### mongodb 存储引擎 (WiredTiger [BTree])
 MongoDB从3.0开始引入可插拔存储引擎的概念。目前主要有MMAPV1、WiredTiger存储引擎可供选择。在3.2版本之前MMAPV1是默认的存储引擎,其采用linux操作系统内存映射技术,但一直饱受诟病；3.4以上版本默认的存储引擎是wiredTiger,相对于MMAPV1其有如下优势:
 1. 读写操作性能更好,WiredTiger能更好的发挥多核系统的处理能力；
 2. MMAPV1引擎使用表级锁,当某个单表上有并发的操作,吞吐将受到限制。WiredTiger使用文档级锁,由此带来并发及吞吐的提高
 3. 提供压缩算法,可以大大降低对硬盘资源的消耗,节省约60%以上的硬盘资源；
 ### mongodb 的基本操作
 #### 新增数据
 ```sql
insert into test.user (name, age ) values('jack', 23);
```
```js
db.user.insertOne({name:"jack",age:23});
db.user.insertMany([{name:"jack",age:23},{name:"mark",age:15}]);
```
```go
s, err := mgo.Dial("mongodb://localhost:27017/test")
if err !=nil {
   panic(err)
}
ms := s.Copy()
defer ms.Close()
c := ms.DB("test").C("user")
doc := bson.M{
 "name":"jack",
 "age":23,
}
c.Insert(doc)

d1:= bson.M{
    "name":"jack",
    "age":23,
   }
d2:=bson.M{
     "name":"mark",
     "age":15,
    }
dos = append(dos,d1,d2)
c.Insert(d1,d2)
```
 #### 更新数据
 ```sql
update test.user set age =18 where name ='jack';
```
```js
db.user.updateOne({name:"jack"},{"$set":{age:18}});
db.user.updateMany({name:"jack"},{"$set":{age:18}});
```
```go
c := ms.DB("test").C("user")
selector:=bson.M{
    "name":"jack",
}
opt :=bson.M{
    "$set":bson.M{
        "age":18,
   },
}
c.Update(selector,opt)

c.UpdaeAll(selector,opt)
```
#### 删除数据
```sql
delete from test.user where name ='jack';
```
```js
db.user.deleteOne({name:"jack"}); // 删除找到的第一条
db.user.deleteMany({name:"jack"}); // 删除所有
```
```go
c := ms.DB("test").C("user")
selector:=bson.M{
    "name":"jack",
}
c.Remove(selector)

c.RemoveAll(selector)
```
#### 查询数据

```sql
select  * from test.user where name ='jack' limit 1,2;
```

```js
db.user.find({name:"jack"}).skip(1).limit(2);
```

```go
query := bson.M{
    "name":"jack",
}
res:=bson.M{}
c.Find(query).One(res)
allRes :=[]bson.M{}
c.Find(query).Skip(1).Limit(2).All(&allRes)
```
#### 聚合查询

```sql
select
	dept ,
	sum(score) total_score,
	avg(score) avg_score
from
	test.`user`
where
	id < 100
GROUP by
	dept
	limit 0,2;
```

```js
db.user.aggregate([
    {
        $match:{
            _id :{
            $gt:100
            }
        }
    },
    {
        $group:{
            _id:"$dept",
            sum_score :{$sum:"$score"},
            avg_score: { $avg:"$score" },
            }
    },
    {
        $skip:0,
    },
    {
        $limit:2,
    }
]);
```
```go
pipeline:=[]bson.M{
    bson.M{
    "$match":bson.M{
        "_id":bson.M{
                "$gt":100,
             },
          },
    },
    bson.M{
    "$group":bson.M{
        "_id":"$dept",
        "sum_score":bson.M{
            "$sum":"$score",
         },
        "$avg_score":bson.M{
            "$avg": "$score",
        },
    },
  },
    bson.M{
    "$skip":0,
  },
  bson.M{
     "$limit":2,
   }
}
allRes :=make([]bson.M,0)
c.Pipe(pipeline).All(&allRes)
```

```sql
select u.*,d.dept_name from  user u, dept d where u.dept = d.id;
```

```js
db.user.aggregate([
    {
      $lookup:{
        from:"dept",
        localField:"dept",
        foreignField:"_id",
        as:"deptArr"
      }
    },
    {
        $addFields:{
            "deptObj":{
                  $arrayElemAt:["$deptArr",0],
             }
        }
    },
    {
         $project:{
                "deptArr" :0,
          }
    }
]);
```

```go
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
        c.Pipe("user", pipeline).All(&allRes)
```


