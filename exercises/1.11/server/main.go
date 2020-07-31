package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("%s %s", req.Method, req.URL.Path)
		// sleep forever
		for true {
			time.Sleep(5 * time.Minute)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
