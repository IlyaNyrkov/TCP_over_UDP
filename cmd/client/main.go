package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	p := make([]byte, 2048)
	conn, err := net.Dial("udp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("Some error %v", err)
		panic(err)
	}
	for {
		message := ""
		fmt.Scan(&message)
		fmt.Fprintf(conn, message)
		_, err = bufio.NewReader(conn).Read(p)
		if err == nil {
			fmt.Printf("Answer from server: %s\n", p)
		} else {
			fmt.Printf("Some error %v\n", err)
		}
	}
	conn.Close()
}
