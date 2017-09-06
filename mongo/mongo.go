package main

import (
	"golib/mongo"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name string `bson:"name"`
	Age int `bson:"age"`
}

type MongoInitConfig struct{}

var globalMgoSession mongo.MongoSession

func (_ MongoInitConfig) Get() *mongo.MongodbConfig {
	var cfg = mongo.MongodbConfig{
		User: "root",
		Password: "123",
		Host: "127.0.0.1:27017",
		DbName: "admin",
	}
	return &cfg
}

func mongoConnect() {
	globalMgoSession = mongo.Init(MongoInitConfig{})
}

func main(){
	mongoConnect()
	insert()
}

func insert(){
	_, sess, def := globalMgoSession.Get()
	defer def()

	c := sess.DB("test").C("testCollection")

	changeInfo, err := c.Upsert(nil,bson.M{"name": "newJohn", "age": 110})


	fmt.Println(changeInfo.UpsertedId)

	if err != nil {
		fmt.Println("Error ocured", err)
	}
}
