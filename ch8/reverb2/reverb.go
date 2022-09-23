// see page 224

// reverb2 is a echo server that simulates the reverberations of a real echo.
// This is to demonstrate a server that uses multiple goroutines per connection.
// The addition of the go keyword to the echo function call enables a more realistic
// response.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // .e.g., connection aborted
			continue
		}
		go handleConn(conn) // handles connections concurrently
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second) // go keyword enables goroutines
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

// Requires upgrading the client (netcat) program so that it sends 
// terminal input to the server while also copying the server response
// to the output, which presents another opportunity to use concurrency.