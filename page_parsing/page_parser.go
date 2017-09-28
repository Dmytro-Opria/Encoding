package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nrparser/parser"
)


type Err struct {
	ErrStr string `json:"error"`
}

func main() {
	fmt.Println("start")
	http.HandleFunc("/json", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start")
	url := r.FormValue("url")

	page, err := parser.parsePage(url, true)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {

		errorHandle(w, err)

		return
	}

	byteBuf, err := json.Marshal(page)

	if err != nil {

		errorHandle(w, err)

		return
	}
	w.Write(byteBuf)
}

func errorHandle(w http.ResponseWriter, err error) {

	fmt.Println(err)

	parsedErr := Err{ErrStr: err.Error()}

	json.NewEncoder(w).Encode(parsedErr)
}
