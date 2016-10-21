package main

import (
	"fmt"
	"net"
	"os"
)

func sender(conn net.Conn) {
		words := "hello world!"
		conn.Write([]byte(words))
		buffer := make([]byte, 2048)
		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println("client connection error: ", err)
			return
		}
		fmt.Println("client receive data string:\n", string(buffer[:n]))
		
	    fmt.Println("send over")

}



func main() {
	//server := "127.0.0.1:1024"
	server := "192.168.6.81:1024"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}


	fmt.Println("connect success")
	sender(conn)

}

