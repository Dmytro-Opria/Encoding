package main

import (
	"time"
	"fmt"
)

func main(){
	//tm()
	//epoch()
	timeFormat()
}

func tm() {
	p := fmt.Println
	now := time.Now()
	p(now)
	p(now.Year())
	p(now.Month())
	p(now.Day())
	p(now.Hour())
	p(now.Minute())
	p(now.Second())
	p(now.Location())
	//loc, _ := time.LoadLocation("Europe/London")

	//set timezone,
	then := time.Date(2017, 6, 7, 11, 0,0,0, time.Now().Location())

	p((int(now.Sub(then).Hours()) + 14 * 24)/24)
}

func epoch() {
	p := fmt.Println

	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()

	p(secs/60/60/24/365)
	p(now.UnixNano())

	p(time.Unix(secs, 0))
	p(time.Unix(secs, 60*5*1000000000))// 5 minutes difference

	p(time.Unix(0, nanos))
	p(time.Unix(60 * 5, nanos))// 5 minutes difference
}

func timeFormat() {
	p := fmt.Println
	now := time.Now()
	p(now.Format(time.Kitchen))
	p(now.Format("Day: Mon\nMonth: Jan\nTime: 15:04:05"))
}