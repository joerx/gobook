package main

import (
	"fmt"
	"io"
	"strings"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {

		if ! strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		fmt.Fprintf(os.Stdout, "fetch %s\n", url)

		// fetch thing
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "%s\n", response.Status)

		// decode body of thing
		io.Copy(os.Stdout, response.Body)
		response.Body.Close()
	}
}
