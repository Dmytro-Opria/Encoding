package main

import (
	"fmt"
	"golib/mongo"
	"gopkg.in/mgo.v2/bson"
	"sync"
	"time"
)

type Person struct {
	Name       string    `bson:"name"`
	Age        int       `bson:"age"`
	LastUpdate time.Time `bson:"lastupdate"`
	Count      int       `bson:"count"`
}
type InvalidBidRequest struct {
	TickersCompositeId int       `bson:"tickersCompositeId"`
	Date               time.Time `bson:"date"`
	Reason             string    `bson:"reason"`
}

var (
	globalMgoSession mongo.MongoSession
	invalidBidMu     sync.Mutex
	invalidBids      = make(map[InvalidBidRequest]int, 0)
)

type MongoInitConfig struct{}

func (_ MongoInitConfig) Get() *mongo.MongodbConfig {
	var cfg = mongo.MongodbConfig{
		User:     "root",
		Password: "123",
		Host:     "127.0.0.1:27017",
		DbName:   "admin",
	}
	return &cfg
}

func mongoConnect() {
	globalMgoSession = mongo.Init(MongoInitConfig{})
}

func main() {
	mongoConnect()
	//insert()
	//fmt.Println(findAndSort())
	//fmt.Println(findFirstOne())
	//fmt.Println(findLastOne())
	//fmt.Println(findWithNatural())
	//incrementAllCounters()
	ToInvalidBidsList(1, time.Now(), []string{"First error"})
	ToInvalidBidsList(1, time.Now(), []string{"First error"})
	ToInvalidBidsList(1, time.Now(), []string{"First error"})
	ToInvalidBidsList(1, time.Now(), []string{"First error"})
	ToInvalidBidsList(2, time.Now(), []string{"Second error"})
	ToInvalidBidsList(3, time.Now(), []string{"Third error"})
	ToInvalidBidsList(4, time.Now(), []string{"Fourth error"})
	ToInvalidBidsList(5, time.Now(), []string{"Five error","Six error"})
	SaveInvalidBidRequests()
}

func insert() {
	_, sess, def := globalMgoSession.Get()
	defer def()

	c := sess.DB("test").C("testCollection")

	changeInfo, err := c.Upsert(bson.M{"_id": bson.NewObjectId()}, bson.M{"name": "Hermiona", "age": 27})

	fmt.Println(changeInfo.UpsertedId)

	if err != nil {
		fmt.Println("Error ocured", err)
	}
}

func findAndSort() (result []Person) {
	_, sess, def := globalMgoSession.Get()
	defer def()

	c := sess.DB("test").C("testCollection")

	c.Find(nil).Sort("age").All(&result)

	return
}

func findFirstOne() (res Person) {
	_, sess, def := globalMgoSession.Get()
	defer def()

	c := sess.DB("test").C("testCollection")

	c.Find(nil).One(&res)

	return
}

func findLastOne() (res Person) {
	_, sess, def := globalMgoSession.Get()
	defer def()

	c := sess.DB("test").C("testCollection")

	dbSize, err := c.Count()
	if err != nil {
		return
	}

	err = c.Find(nil).Skip(dbSize - 1).One(&res)
	if err != nil {
		return
	}

	return
}

func findWithNatural() (res []Person) {
	_, sess, def := globalMgoSession.Get()
	defer def()

	c := sess.DB("test").C("testCollection")

	c.Find(nil).Sort("-$natural").Limit(2).All(&res)

	return
}

func incrementAllCounters() {
	_, sess, def := globalMgoSession.Get()
	defer def()

	c := sess.DB("test").C("testCollection")

	change := bson.M{"$inc": bson.M{"counter": 1}, "$set": bson.M{"lastUpdate": time.Now()}}
	info, err := c.UpdateAll(bson.M{}, change)
	if err != nil {
		fmt.Println("Error ocured", err)
	}
	fmt.Println(info)
}

func ToInvalidBidsList(tickerCompositeId int, tm time.Time, reasons []string) {
	date := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, tm.Location())

	key := InvalidBidRequest{}
	key.TickersCompositeId = tickerCompositeId
	key.Date = date
	for i, reason := range reasons {
		if len(reasons) == i+1 {
			key.Reason += reason
		} else {
			key.Reason += reason + ", "
		}
	}

	if val, ok := invalidBids[key]; ok {
		invalidBids[key] = val + 1
		return
	}
	invalidBids[key] = 1
}

func SaveInvalidBidRequests() {
	saveInvalidBidRequests(invalidBids)
}

func saveInvalidBidRequests(bidRequests map[InvalidBidRequest]int) {

	if len(bidRequests) == 0 {
		return
	}

	invalidBidMu.Lock()
	defer invalidBidMu.Unlock()

	dbname, sess, def := globalMgoSession.Get()
	defer def()

	c := sess.DB(dbname).C("invalid_bid_requests")

	for key, value := range bidRequests {

		if _, err := c.Upsert(key,bson.M{"$inc": bson.M{"counter": value}/*, "$set": key*/}); err != nil {
			fmt.Printf("save invalidBidRequest %s %+v\n", err.Error(), key)
		}

		delete(bidRequests, key)
	}
}
