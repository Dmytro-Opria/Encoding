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
	//insert()
	//fmt.Println(findAndSort())
	//fmt.Println(findFirstOne())
	//fmt.Println(findLastOne())
	fmt.Println(findWithNatural())
}

func insert(){
	_, sess, def := globalMgoSession.Get()
	defer def()

	c := sess.DB("test").C("testCollection")

	changeInfo, err := c.Upsert(bson.M{"_id":bson.NewObjectId()},bson.M{"name": "Hermiona", "age": 27})


	fmt.Println(changeInfo.UpsertedId)

	if err != nil {
		fmt.Println("Error ocured", err)
	}
}

func findAndSort()(result []Person){
	_, sess, def := globalMgoSession.Get()
	defer def()

	c := sess.DB("test").C("testCollection")

	c.Find(nil).Sort("age").All(&result)

	return
}

func findFirstOne()(res Person){
	_, sess, def := globalMgoSession.Get()
	defer def()

	c := sess.DB("test").C("testCollection")

	c.Find(nil).One(&res)

	return
}

func findLastOne()(res Person){
	_, sess, def := globalMgoSession.Get()
	defer def()

	c := sess.DB("test").C("testCollection")

	dbSize, err := c.Count()
	if err != nil {
		return
	}

	err = c.Find(nil).Skip(dbSize-1).One(&res)
	if err != nil {
		return
	}

	return
}

func findWithNatural()(res []Person){
	_, sess, def := globalMgoSession.Get()
	defer def()

	c := sess.DB("test").C("testCollection")


	c.Find(nil).Sort("-$natural").Limit(2).All(&res)

	return
}
