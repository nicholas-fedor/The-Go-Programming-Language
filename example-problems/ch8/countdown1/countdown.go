// see page 244

// Countdown does the countdown for a rocket launch.
// The time.Tick function returns a channel on which it sends events periodically,
// acting like a metronome. The value of each event is a timestamp, but it is
// rarely as interesting as the fact of its delivery.
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

func launch()  {
	fmt.Println("Lift Off!")
}