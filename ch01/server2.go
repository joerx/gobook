package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mx sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handleHello)
	http.HandleFunc("/count", handleCount)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	log.Print("hello")
	mx.Lock()
	count++
	mx.Unlock()
	fmt.Fprintf(w, "Path: %q\n", r.URL.Path)
} 

func handleCount(w http.ResponseWriter, r *http.Request) {
	log.Print("count")
	mx.Lock()
	fmt.Fprintf(w, "Count: %d\n", count)
	mx.Unlock()
}
