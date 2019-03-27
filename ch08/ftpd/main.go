package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

// ftp session
// 220 FTP server ready
// > user <username>
// 331 Please specify the password.
// > pass <password>
// 230 Login successful.
// > pasv
// 227 Entering Passive Mode (4,31,198,44,255,203) -- open new connection given port, send port number int(n/256), n%256
// > list
// 150 Here comes the directory listing. -- send via the previously opened data connection
// 226 Directory send OK.
// > retr
// ...

const user = "files"
const pass = "files"
const port = 2100

func main() {
	// slice := []string{"one"}
	// fmt.Println(slice[0], slice[1:])
	if err := ftpServer(); err != nil {
		log.Fatal(err)
	}
}

func ftpServer() error {

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return err
	}

	log.Printf("Ready, will now accept connections on port %d", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go func() {
			s := &session{}
			defer conn.Close()
			if err := s.handleConn(conn); err != nil {
				log.Println(err)
			}
		}()
	}
}

type session struct {
	// conn net.Conn
	// rw *bufio.ReadWriter
}

// func newSession(conn net.Conn) *session {
// 	return &session{rw: bufio.NewReadWriter(
// 		bufio.NewReader(conn),
// 		bufio.NewWriter(conn),
// 	)}
// }

func readCommand(r *bufio.Reader) (string, []string, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return "", nil, err
	}

	line = strings.TrimSuffix(line, "\r\n")
	parts := strings.Split(line, " ")

	return strings.ToLower(parts[0]), parts[1:], nil
}

func sendResponse(w io.Writer, code int, message string) error {
	if _, err := fmt.Fprintf(w, "%3d %s\n", code, message); err != nil {
		return err
	}
	return nil
}

func mustLogin(ctrl net.Conn) error {
	var user string
	var pass string

	r := bufio.NewReader(ctrl)

	// keep user in login loop until the provide username and password
	for user == "" || pass == "" {
		cmd, args, err := readCommand(r)
		if err != nil {
			return err
		}
		switch cmd {
		case "user":
			user = args[0]
			sendResponse(ctrl, 331, "Please specify password")
		case "pass":
			pass = args[0]
		default:
			sendResponse(ctrl, 530, "Please login with USER and PASS.")
		}
	}

	sendResponse(ctrl, 230, "Login successful.")
	return nil
}

func (s *session) handleConn(conn net.Conn) error {
	if err := sendResponse(conn, 220, "aWes0meFTP 0.0.1-alpha is ready"); err != nil {
		return err
	}

	// mustLogin(conn)

	r := bufio.NewReader(conn)

	var dataConn net.Conn

	for {
		cmd, _, err := readCommand(r)
		if err != nil {
			break
		}

		log.Printf("Received '%s'", cmd)

		switch cmd {
		case "pwd":
			err = sendPWD(conn)
		case "list":
			err = sendList(dataConn)
			dataConn.Close()
			dataConn = nil
		case "pasv":
			dataConn, err = openDataConn(conn)
			go (func() {
				<-time.NewTimer(time.Second * 10).C
				log.Println("data channel timeout")
				dataConn.Close()
				dataConn = nil
			})()
		default:
			err = sendResponse(conn, 530, "Unknown command")
		}

		if err != nil {
			break
		}
	}

	return nil
}

func sendPWD(dc net.Conn) error {
	return sendResponse(dc, 220, "/")
}

func sendList(dc net.Conn) error {
	return sendResponse(dc, 220, "... a list of files ...")
}

// Send data connection address to client, then wait for them to accept connection
func openDataConn(ctrl net.Conn) (net.Conn, error) {
	port := 10900
	i1 := int(port / 256)
	i2 := port % 256

	log.Printf("Data channel on port %d", port)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	err = sendResponse(ctrl, 220, fmt.Sprintf("(127,0,0,1,%d,%d)", i1, i2))
	if err != nil {
		l.Close()
		return nil, err
	}

	return l.Accept()
}
