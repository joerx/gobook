package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type fetchOpts struct {
	logFile string
	urls    []string
}

func main() {
	opts := getopts()
	out, err := getoutput(opts)
	if err != nil {
		fatal(err)
	}
	fetchAll(opts.urls, out)
}

func fatal(err error) {
	fmt.Fprint(os.Stderr, err)
	os.Exit(1)
}

func getoutput(opts fetchOpts) (io.Writer, error) {
	if opts.logFile != "" {
		w, err := os.Create(opts.logFile)
		if err != nil {
			return nil, err
		}
		return w, nil
	}
	return os.Stdout, nil
}

func getopts() (opts fetchOpts) {
	flag.StringVar(&opts.logFile, "log", "", "File to write output to")
	flag.Parse()
	opts.urls = flag.Args()
	return opts
}

func fetchAll(urls []string, out io.Writer) error {
	c := make(chan error)

	for _, url := range urls {
		go fetch(url, out, c)
	}

	// block until all urls are fetched
	for range urls {
		if err := <-c; err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		}
	}

	return nil
}

func fetch(url string, out io.Writer, c chan error) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		c <- err
		return
	}

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
