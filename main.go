package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: rauxy <command> [arguments]")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) != 4 {
			log.Fatalln("Usage: rauxy add <token_name> <port>")
		}
		tokenName := os.Args[2]
		port := os.Args[3]
		err := AddToken(tokenName, port)
		if err != nil {
			log.Fatalf("Error adding token: %v\n", err)
		}
		fmt.Println("Token added successfully.")
	case "rm":
		if len(os.Args) != 3 {
			log.Fatalln("Usage: rauxy rm <token_name>")
		}
		tokenName := os.Args[2]
		err := RemoveToken(tokenName)
		if err != nil {
			log.Fatalf("Error removing token: %v\n", err)
		}
		fmt.Println("Token removed successfully.")
	case "ls":
		err := ListTokens()
		if err != nil {
			log.Fatalf("Error listing tokens: %v\n", err)
		}
	case "serve":
		if len(os.Args) != 4 {
			log.Fatalln("Usage: rauxy serve <port_source> <port_dest>")
		}
		portSource := os.Args[2]
		portDest := os.Args[3]
		err := Serve(portSource, portDest)
		if err != nil {
			log.Fatalf("Error starting server: %v\n", err)
		}
	default:
		fmt.Println("Unknown command:", command)
	}
}
