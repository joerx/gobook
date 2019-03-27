package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	addr := flag.String("a", ":8000", "Address")
	mode := flag.String("m", "c", "Operation mode, (s)erver or (c)lient")

	flag.Parse()

	switch *mode {
	case "s":
		server(*addr)
	case "c":
		client(*addr)
	default:
		fmt.Fprintf(os.Stderr, "Invalid mode: %s must be \"s\" or \"c\"\n", *mode)
		os.Exit(1)
	}
}

func client(addr string) {
	log.Printf("Connect to %s", addr)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	for {
		t, err := input("> ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(conn, t)
		if err := response(conn); err != nil {
			fmt.Println(err)
			break
		}
	}
}

func server(addr string) {

}

func input(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s ", prompt)
	return reader.ReadString('\n')
}

func response(conn net.Conn) error {
	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return err
		}
		fmt.Print(line)
	}
}
