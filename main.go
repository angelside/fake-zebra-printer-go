/*
The startServer function starts a server that listens for incoming connections on IP address 127.0.0.1 and port 9100. When a connection is accepted, the handleConnection function is called in a separate goroutine to handle the incoming data. The handleConnection function reads lines of ZPL code from the connection using a scanner, and for each line of input it calls the Print method of the Printer struct to get the response. The response is then written back to the connection using the fmt.Fprintln
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	//"strings"
    //"regexp"
)

func main() {
	startServer()
}

type Printer struct{}

// Sends a string of ZPL code to the printer and returns the printer's response
func (p *Printer) Print(zpl string) string {

	// TODO: later remove it
	fmt.Println(zpl)
	return zpl

	// TODO: Catch password reset
	// ^XA^KP%s^JUS^XZ


	/*
	// use a regular expression to search for the pattern ^FD followed by a sequence of digits
	re := regexp.MustCompile(`\^FD(\d+)`)
	match := re.FindStringSubmatch(zpl)
	if len(match) > 1 {
		// the first group of the regular expression is the sequence of digits
		fmt.Println(match[1]) // prints "105925"
	}
	*/

	/*
	// parse the ZPL code and perform the corresponding actions
	if strings.Contains(zpl, "^FO50,50") {
		// ^FO50,50 indicates the start of a text field at position 50,50
		return "OK"
	} else if strings.Contains(zpl, "^FDHello, World!^FS") {
		// ^FD specifies the field data and ^FS indicates the end of the field
		return "PRINTED: Hello, World!"
	}
	return "INVALID ZPL CODE"
	*/
}

func startServer() {
	p := &Printer{}

	// Listen for incoming connections on IP address 127.0.0.1 and port 9100
	ln, err := net.Listen("tcp", "127.0.0.1:9100")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	fmt.Println("Listening for connections on 127.0.0.1:9100")

	// Accept incoming connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn, p)
	}
}

// Reads ZPL code from the connection and sends the response back to the client
func handleConnection(conn net.Conn, p *Printer) {
	// Create a scanner to read from the connection
	scanner := bufio.NewScanner(conn)

	// Reads lines of ZPL code from the connection
	for scanner.Scan() {
		zpl := scanner.Text()
		response := p.Print(zpl)
		fmt.Fprintln(conn, response)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	conn.Close()
}
