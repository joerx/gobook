package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	fetchAll(os.Args[1:], os.Stdout)
	fmt.Fprintln(os.Stderr, "All done")
}

func fatal(err error) {
	fmt.Fprint(os.Stderr, err)
	os.Exit(1)
}

func fetchAll(urls []string, out io.Writer) error {
	c := make(chan error)

	client := &http.Client{}
	defer client.CloseIdleConnections()

	for _, url := range urls {
		go fetch(url, client, out, c)
	}

	// block until all urls are fetched
	for range urls {
		if err := <-c; err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		}
	}

	return nil
}

func fetch(url string, client *http.Client, out io.Writer, c chan error) {
	start := time.Now()

	abort := make(chan struct{})
	pass := make(chan struct{})

	timer := time.NewTimer(5 * time.Second)

	// circuit breaker
	go func() {
		select {
		case <-timer.C:
			close(abort)
		case <-pass:
			timer.Stop()
		}
	}()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("%#v\n", err)
		c <- err
		return
	}
	req.Cancel = abort
	req.Close = true

	resp, err := client.Do(req)
	if err != nil {
		c <- err
		return
	}

	close(pass) // we got a response, cancel circuit breaker

	defer resp.Body.Close()
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		c <- err
		return
	}

	msecs := time.Since(start).Milliseconds()
	fmt.Fprintf(out, "GET %s %d read %d bytes in %d ms\n", url, resp.StatusCode, nbytes, msecs)

	c <- nil
}
