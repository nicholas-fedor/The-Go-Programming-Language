// see page 221

// Netcat1 is a read-only TCP client.
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
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader)  {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

// Explanation:
// This program reads data from the connection and writes it to the standard output
// until an end-of-file condition or error occurs.
// The mustCopy function is a utility used in several examples in this section.

// The second client must wait until the first client is finished because the server
// is sequential; it deals with only one client at a time.
// Just one small change is needed to make the server concurrent: 
// adding the go keyword to the call to handleConn causes each call to run 
// in its own goroutine.

// Usage:
// go build gopl.io/ch8/netcat1
// Run in two terminals while the clock1 program is running.
// ./netcat1
// To kill the clock1 program (if using the & flag, use the command: killall clock1)