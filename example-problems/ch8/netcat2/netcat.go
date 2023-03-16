// see page 223

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn) // enables concurrent calls to the server.
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader)  {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

// Explanation:
// While the main goroutine reads the standard input and sends it to the server, 
// a second goroutine reads and prints the server's response.
// When the main goroutine encounters the end of the input, for example, after 
// the user types Control-D (^D) at the terminal (or Control-Z for Windows), 
// the program stops, even if the other goroutine still has work to do.

// Usage: 
// Build and run both (reverb1 and netcat2) programs. 

// Input:	Hello?
// Output:			HELLO?
// Output:			Hello?
// Output:			hello?
// Input:	Is there anyone there?
// Output:			IS THERE ANYONE THERE?
// Output:			Is there anyone there?
// Output:			is there anyone there?
// Input:	Yoooo-hooo!
// Output:			YOOOO-HOOO!
// Output:			Yoooo-hooo!
// Output:			yoooo-hooo!