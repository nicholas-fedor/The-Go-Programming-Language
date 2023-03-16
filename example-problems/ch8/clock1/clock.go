// see page 219

// Clock1 is a TCP server that periodically writes the current time
// to the client once per second.
package main

import (
	"io"
	"log"
	"net"
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
			log.Print(err) // e.g., connection aborted
			continue
		}
		handleConn(conn) // handle one connection at a time.
	}
}

func handleConn(c net.Conn)  {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

// Explanation:
// The Listen function creates a net.Listener, an object that listens for
// incoming connections on a network port, in this case TCP port
// localhost:8000.
// The listener's Accept method blocks until an incoming connection
// request is made, then returns a net.Conn object representing the connection.

// The handleConn function handles one complete client connection. In a loop, it 
// writes the current time, time.Now(), to the client. Since net.Conn satisfies
// the io.Writer interface, we can write directly to it. The loop ends when the
// write fails, most likely because the client has disconnected, at which point
// handleConn closes its side of the connection using a deferred call to Close
// and goes back to waiting for another connection request.

// The time.Time.Format method provides a way to format date and time information
// by example. Its argument is a template indicating how to format a reference
// time, specifically Mon Jan 2 03:04:05PM 2006 UTC-0700. The reference time has 
// eight components (day of the week, month, day of the month, and so on).
// Any collection of them can appear in the Format string in any order and in 
// a number of formats; the selected components of the date and time will be 
// displayed in the selected formats.
// Here we are just using the hour, minute, and second of the time.
// The time package defines templates for many standard time formats, such as 
// time.RFC1123. The same mechanism is used in reverse when parsing a time using 
// time.Parse.

// Usage: 
// go build gopl.io/ch8/clock1
// ./clock1 & 
// *the & flag allows the program to run without hogging the terminal.
// nc localhost 8000

// Output:
// 13:58:54
// 13:58:55
// 13:58:56
// 13:58:57
