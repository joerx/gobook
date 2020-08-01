package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", echoHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]: %q\n", k, v)
	}

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "RemoteAddr: %q\n", r.RemoteAddr)
	fmt.Fprintf(w, "\n")

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q]: %q\n", k, v)
	}
}
