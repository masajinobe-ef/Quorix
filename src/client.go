package src

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func StartClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:3333")
	if err != nil {
		fmt.Println("Error of Connection to Server:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to Server!")

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println("Server:", scanner.Text())
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		_, err := conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Error of Sending message:", err)
			return
		}
	}
}
