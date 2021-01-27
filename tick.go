package main

import (
	"flag"
	"fmt"
	"time"
)

// Runs a timer for the specified number of minutes and seconds.
// For example, to run a timer for 1 minute and 30 seconds:
//
//     tick -m 1 -s 30
func main() {
	seconds := flag.Int("s", 0, "Number of seconds the timer should run")
	minutes := flag.Int("m", 0, "Number of minutes the timer should run")
	flag.Parse()

	duration := (time.Duration(*seconds) * time.Second) + (time.Duration(*minutes) * time.Minute)
	ding := time.After(duration)
	tick := time.Tick(1 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Printf("\033[2K\r%s", duration)
			duration = duration - (1 * time.Second)
		case <-ding:
			fmt.Printf("\033[2K\r%s", duration)
			fmt.Println("\nDING!")
			return
		}
	}
}
