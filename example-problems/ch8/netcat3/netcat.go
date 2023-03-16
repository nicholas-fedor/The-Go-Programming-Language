// see page 227

// netcat3 demonstrates the use of unbuffered channels to synchronize two goroutines.
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
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

// Explanation:
// When the user closes the standard input stream, mustCopy returns the main goroutine
// calls conn.Close(), closing both halves of the network connection. Closing the write half of
// the connection causes the server to see an end-of-file condition. Closing the read half causes
// the background goroutine's call to io.Copy to return a "read from closed connection" error,
// which is why we've removed the error logging.
// * Notice that the go statement calls a literal function.

// Before it returns, the background goroutine logs a message, then sends a value on the done
// channel. The main goroutine waits until it has received this value before returning. As a
// result, the program always logs the "done" message before exiting.

// Messages sent over channels have two important aspects. Each message has a value, but
// sometimes the fact of communication and the moment at which it occurs are just as
// important. We call messages events when we wish to stress this aspect. When the event carries
// no additional information, that is, its sole purpose is synchronization, we'll emphasize this
// by using a channel whose element type is struct{}, though it's common to use a channel of
// bool or int for the same purpose since done <- 1 is shorter than done <- struct{}{}.
