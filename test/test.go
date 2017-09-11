package main

import (
	"os"
	"encoding/json"
	p "gossp/pool"
)
type VrItem struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Img      string `json:"img"`
	Link     string `json:"link"`
	Hash     string `json:"hash"`
	VideoExt string `json:"videoExt"`
}

type Test struct {
	Debug   int      `json:"debug"`
	Version string   `json:"version"`
	Method  string   `json:"method"`
	Device  string   `json:"device"`
	Auction struct{} `json:"auction"`
	Adblock struct {
		Total int `json:"total"`
		Size  int `json:"size"`
	} `json:"adblock"`
	Parameters struct {
		SourceId  uint32 `json:"source id"`
		Page      uint8  `json:"page"`
		Subnet    int    `json:"subnet"`
		Os        int    `json:"os"`
		Browser   int    `json:"browser"`
		Protocol  int    `json:"protocol"`
		Ip        string `json:"ip"`
		UserAgent string `json:"user-agent"`
		Adblock struct {
			Etalon int `json:"etalon"`
			Offset int `json:"offset"`
			Id     int `json:"id"`
		} `json:"adblock"`
		Region struct {
			Country int `json:"country"`
			Maxmind int `json:"maxmind"`
			Id      int `json:"id"`
		} `json:"region"`
		Auditories []int `json:"auditories"`
		Interests  []int `json:"interests"`
	} `json:"parameters"`
	Update struct {
		Current int32 `json:"current"`
	} `json:"update"`
	InformerComposite struct {
		Awc           struct{} `json:"awc"`
		DeviceType    string   `json:"dt"`
		TrafficSource string   `json:"ts"`
		TrafficType   string   `json:"tt"`
		IsBot         int      `json:"isBot"`
		H2            string   `json:"h2"`
		Config struct {
			Mute            string `json:"mute"`
			Capping         int    `json:"capping"`
			CappingInterval int    `json:"cappingInterval"`
			Uid             int64  `json:"uid"`
			Vast            string `json:"vast"`
			WagesTypes      string `json:"wages_types"`
		} `json:"config"`
		Template   string   `json:"template"`
		Styles     string   `json:"styles"`
		Lib        string   `json:"lib"`
		VRplaylist []VrItem `json:"vr_playlist"`
	}
}

func main() {

	outBuf := p.GetBuffer()

	enc := json.NewEncoder(outBuf)

	combInfo := Test{}

	combInfo.InformerComposite.Awc = struct{}{}
	combInfo.InformerComposite.DeviceType = "desktop"
	combInfo.InformerComposite.TrafficSource = ""
	combInfo.InformerComposite.TrafficType = "Direct"
	combInfo.InformerComposite.IsBot = 0
	combInfo.InformerComposite.H2 = "Fi9njqm0PrqufudmZ9OHMA"

	combInfo.InformerComposite.Config.Mute = "1"
	combInfo.InformerComposite.Config.Capping = 0
	combInfo.InformerComposite.Config.CappingInterval = 0
	combInfo.InformerComposite.Config.Uid = 5535287
	combInfo.InformerComposite.Config.Vast = "//servicer.mgid.com/108978/?vast=1"
	combInfo.InformerComposite.Config.WagesTypes = "video"

	combInfo.InformerComposite.Template = "<div class=\"mgbox\">..."
	combInfo.InformerComposite.Styles = ".mgVideoBottom_89716 {\r\n    position: absolute;..."
	combInfo.InformerComposite.Lib = "4"

	playElem1 := VrItem{}
	playElem1.Id =  "abcd1"
	playElem1.Title = "text1"
	playElem1.Img = "image1-url.com"
	playElem1.Link = "//video-cdn.mgid.com/file1.mp4"
	playElem1.Hash = "1"
	playElem1.VideoExt = "eyJkdXJhdGlvbiI6Mzd9"

	combInfo.InformerComposite.VRplaylist = append(combInfo.InformerComposite.VRplaylist, playElem1)

	playElem2 := VrItem{}

	playElem2.Id = "abcd2"
	playElem2.Title = "text2"
	playElem2.Img = "image2-url.com"
	playElem2.Link = "//video-cdn.mgid.com/file2.mp4"
	playElem2.Hash = "1"
	playElem2.VideoExt = "eyJkdXJhdGlvbiI6Mzd9"

	combInfo.InformerComposite.VRplaylist = append(combInfo.InformerComposite.VRplaylist, playElem2)


	//============================


	file, err := os.Create("test.json")
	if err != nil {
		return
	}
	defer file.Close()

	enc.Encode(combInfo)

	file.Write(outBuf.Bytes())
}
