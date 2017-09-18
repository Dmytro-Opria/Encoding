package main
import (
	"fmt"
	"net/http"
	_"golang.org/x/net/http2"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
		if pusher, ok := w.(http.Pusher); ok {
			fmt.Println("push is OK")
			options := &http.PushOptions{
				Header: http.Header{
					"Accept-Encoding": r.Header["Accept-Encoding"],
				},
			}
			if err := pusher.Push("/style.css", options); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		} else {
			fmt.Println("have not push")
		}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("start http listening :3002")
	err := http.ListenAndServeTLS(":3002", "apache2.crt", "apache2.key", nil)
	fmt.Println(err)
}
