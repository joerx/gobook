package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Run as clockwall server",
			Action:  server,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "address, a",
					Value: ":8000",
					Usage: "host and port to listen on",
				},
				cli.StringFlag{
					Name:  "label, l",
					Usage: "label to identify the server",
				},
			},
		},
		{
			Name:    "client",
			Aliases: []string{"c"},
			Usage:   "Run as clockwall client",
			Action:  client,
			Flags: []cli.Flag{
				cli.StringSliceFlag{
					Name:  "address, a",
					Usage: "List of host/ports to listen on",
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func server(c *cli.Context) error {
	fmt.Println("Clockwall server")

	address := c.String("address")
	label := c.String("label")

	server := &ClockServer{label}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go server.HandleConn(conn)
	}
}

// ClockServer represents a simple, labelled server sending the current time to connected clients
// along with its label
type ClockServer struct {
	label string
}

// HandleConn handles an incoming connection to the timeserver, sending the current data to the
// given channel each second.
func (s *ClockServer) HandleConn(c net.Conn) {
	defer c.Close()
	for {
		tstring := time.Now().Format("15:04:05")
		data := fmt.Sprintf("%s=%s\n", s.label, tstring)
		_, err := io.WriteString(c, data)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func client(c *cli.Context) error {
	fmt.Println("Clockwall client")

	addresses := c.StringSlice("address")

	for _, address := range addresses {
		conn, err := net.Dial("tcp", address)
		if err != nil {
			return err
		}
		go mustCopyAndClose(os.Stdout, conn)
	}

	for {
		time.Sleep(1 * time.Minute)
	}
}

func mustCopyAndClose(src io.Writer, conn io.ReadCloser) {
	defer conn.Close()
	if _, err := io.Copy(os.Stdout, conn); err != nil {
		log.Fatal(err)
	}
}
