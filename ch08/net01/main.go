package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// a low level http client
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(conn, "GET / HTTP/1.0\r\n\r\n")

	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
	}
}
