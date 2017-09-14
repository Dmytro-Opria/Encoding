package main

import (
	"strings"
	"fmt"
	"golib/utils"
)

func main() {
	str := `{"mute":"2","wages_types":"video_content","skipoffset":5,"capping":1,"cappingInterval":3,"extLibraries":{"jQuery":{"url":"//code.jquery.com/jquery-3.1.1.min.js","exists":"window.jQuery"},"mgPerfectScroll":{"dep":["jQuery"],"url":"//cdn.mgid.com/js/perfect-scrollbar.js","exists":"window.jQuery && typeof window.Ps !== 'undefined'"}}}  `
	str2 := `{"{\"mute\":\"1\",\"capping\":6,\"cappingInterval\":8,\"extLibraries\":{\"jQuery\":{\"url\":\"//code.jquery.com/jquery-3.1.1.min.js\",\"exists\":\"window.jQuery\"},\"mgPerfectScroll\":{\"dep\":[\"jQuery\"],\"url\":\"//cdn.mgid.com/js/perfect-scrollbar.js\",\"exists\":\"window.jQuery\"}" %!q(bool=true)}`
	str = strings.TrimSuffix(strings.TrimPrefix(str, "{"), "}")

	fmt.Println(getVideoCappingData(str2))

}

func getVideoCappingData(videoCfg string) (mute string, capping, cappingInterval int) {
	strArr := strings.Split(videoCfg, ",")
	for _, val := range strArr {
		if strings.Contains(val, `"mute\"`) || strings.Contains(val, `"mute"`){
			mute = strings.TrimSuffix(strings.TrimPrefix(strings.Split(val, ":")[1], `\"`), `\"`)
			mute = strings.TrimSuffix(strings.TrimPrefix(mute, `"`),`"`)
		}
		if strings.Contains(val, `"capping\"`) || strings.Contains(val, `"capping"`){
			capping = utils.ToInt(strings.Split(val, ":")[1])
		}
		if strings.Contains(val, `"cappingInterval\"`) || strings.Contains(val, `"cappingInterval"`){
			cappingInterval = utils.ToInt(strings.Split(val, ":")[1])
		}
	}
	if mute == "" {
		mute = "0"
	}
	return
}