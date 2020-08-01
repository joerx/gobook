package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/liss", lissHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func lissHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s", req.Method, req.URL.Path)

	params, err := parseParams(req)
	if err != nil {
		badRequest(w, err)
		return
	}

	req.Header.Set("Content-type", "image/gif")
	lissajous(w, params)
}

func badRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-type", "text-plain")
	fmt.Fprintf(w, "bad request - %s", err)
}

func parseParams(req *http.Request) (lissParams, error) {
	var p lissParams

	err := intParam(req.URL.Query(), "cycles", &p.cycles, 5)
	if err != nil {
		return p, err
	}
	err = intParam(req.URL.Query(), "size", &p.size, 100)
	if err != nil {
		return p, err
	}
	err = intParam(req.URL.Query(), "nframes", &p.nframes, 64)
	if err != nil {
		return p, err
	}
	err = intParam(req.URL.Query(), "delay", &p.delay, 8)
	if err != nil {
		return p, err
	}

	return p, nil
}

func intParam(m url.Values, name string, val *int, def int) error {
	strVal := m.Get(name)
	if strVal == "" {
		*val = def
		return nil
	}

	x, err := strconv.Atoi(strVal)
	if err != nil {
		return err
	}

	*val = x
	return nil
}
