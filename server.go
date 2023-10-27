package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"strings"

	"github.com/fatih/color"
)

func startServer(ip, port string) {
	address := fmt.Sprintf("%s:%s", ip, port)

	// Create a TCP listener
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}
	defer listener.Close()

	// Colorize the address
	addressColor := color.New(color.FgCyan).Add(color.Bold)
	address = addressColor.Sprint(address)

	fmt.Printf("Listening for connections on %s \n\n", address)

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	// Create a buffer to accumulate the response
	var responseBuffer bytes.Buffer

	// Reads lines of data from the connection and accumulates them
	for scanner.Scan() {
		code := scanner.Text()
		responseBuffer.WriteString(code)
		responseBuffer.WriteString("\n") // Add a newline between data responses if needed
	}

	// Get the entire response as a single string
	fullResponse := responseBuffer.String()

	// Now you can process the entire response as needed
	processedResponse := processData(fullResponse)

	// Print the processed response to the connection
	fmt.Fprint(conn, processedResponse)
}

// Data processing logic
func processData(response string) string {
	response = strings.TrimSpace(response)

	printData(response)

	return response
}
