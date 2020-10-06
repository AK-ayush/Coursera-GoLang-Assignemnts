package main

import (
	"fmt"
	"time"
)

/* Explaination:
* --------------
* At line 28, when Goroutine 1 is launched, it starts to incrementing
* the variable `count` inside a for loop, meanwhile at line 38, Gorountine 2 is also launched and
* print the current value of `count` variable. As, we can see here that these both Goroutines
* are communicating on variable `count`, so, both the Goroutines's instructions are interleaved.
* And, this is requirement for `Race-Condition`, since the execcution path is non-deterministic.
*
* So, everytime we run this program, this will print the different value of `count`.
***/

func incr(count *int) {
	for i := 0; i < 1000000; i++ {
		(*count)++
	}
}

func print(count int) {
	fmt.Println(count)
}

func main() {
	count := 0

	// Routine 1
	go incr(&count)

	// adding this sleep time, so that above goroutine get some headstart.
	time.Sleep(1 * time.Millisecond)

	// Routine 2
	go print(count)

	// adding sleep time, so that the main thread wait for the
	// go routines to finish.
	time.Sleep(100 * time.Millisecond)
}

/*
In order to see the behaviour, please tun multiple time.
*/
