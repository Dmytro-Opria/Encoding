package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	 _ "net/http/pprof"
)

type Resp struct {
	Stat string `json:"stat"`
}

func handler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(200)

	jsonResp, err := json.Marshal(Resp{"OK"})
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)

}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	jsonResp, err := json.Marshal(Resp{"OK, test"})
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func main() {
	http.HandleFunc("/",handler)
	http.HandleFunc("/test", testHandler)
	http.ListenAndServe(":3000",nil)
}