package main

import (
	"fmt"
	"os"

	"Quorix/src"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Using: go run main.go [server|client]")
		return
	}

	switch os.Args[1] {
	case "server":
		src.StartServer()
	case "client":
		src.StartClient()
	default:
		fmt.Println("Unknown command. Use 'server' or 'client'.")
	}
}
