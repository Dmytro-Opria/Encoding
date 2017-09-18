package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

type Resp struct {
	Stat string `json:"stat"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	jsonResp, err := json.Marshal(Resp{"OK"})
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func pushHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, indexHTML)

	if pusher, ok := w.(http.Pusher); ok {
		options := &http.PushOptions{
			Header: http.Header{
				"Accept-Encoding": r.Header["Accept-Encoding"],
			},
		}
		if err := pusher.Push("/style.css", options); err != nil {
			log.Printf("Failed to push: %v", err)
		}
		if err := pusher.Push("/app.js", options); err != nil {
			log.Printf("Failed to push: %v", err)
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/push", pushHandler)
	fmt.Println("start http listening :3002")
	err := http.ListenAndServeTLS(":3002", "apache2.crt", "apache2.key", nil)
	fmt.Println(err)
}

const indexHTML = `<html>
<head>
	<title>Hello World</title>
	<script src="/app.js"></script>
	<link rel="stylesheet" href="/style.css"">
</head>
<body>
Hey there!
</body>
</html>
`
