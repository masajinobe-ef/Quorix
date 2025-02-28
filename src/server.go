package src

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New connection:", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println("Received:", message)
	}
}

func StartServer() {
	listener, err := net.Listen("tcp", ":3333")
	if err != nil {
		fmt.Println("Error of Server starting:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server started! Port: 3333")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error of Connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
