package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if err := fetchAll(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func fetchAll(urls []string) error {
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("error making http request: %s", err)
		}

		defer resp.Body.Close()
		if _, err = io.Copy(os.Stdout, resp.Body); err != nil {
			return fmt.Errorf("error reading response body: %s", err)
		}
	}

	return nil
}
