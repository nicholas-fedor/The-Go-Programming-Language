// See page 295

// Cross prints the operating system and architecture for which it was built.
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}