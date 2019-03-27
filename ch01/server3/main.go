package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", echoHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%s]: %v\n", k, v)
	}

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "\n")

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%s]: %v\n", k, v)
	}
}
