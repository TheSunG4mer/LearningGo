package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var conn net.Conn

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Ip:port: ")
	ip, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	conn, _ = net.Dial("tcp", ip) // use 127.0.0.1:18081
	defer conn.Close()
	connReader := bufio.NewReader(conn)

	for {
		fmt.Print("> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		if text == "quit\n" {
			return
		}

		fmt.Fprintf(conn, text)
		msg, err := connReader.ReadString('\n')
		if err != nil {
			return
		}

		// Windows uses \r\n as a return character, Trimspace removes the extra '\r' character
		msg = strings.TrimSpace(msg)
		fmt.Println("From server: " + msg)
	}
}
