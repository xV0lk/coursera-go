package main

import (
	"fmt"
	"time"
)

var counter int

// add is a goroutine that increments the counter variable by 1
func add() {
	for i := 0; i < 5; i++ {
		counter++
		fmt.Printf("Incrementing counter to %d\n", counter)
		time.Sleep(time.Millisecond)
	}
}

// print is a goroutine that prints the counter variable
func print() {
	for i := 0; i < 5; i++ {
		fmt.Printf("counter: %d\n", counter)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	// a race condition occurs with the following two goroutines
	// because they are both accessing the same variable
	// and the order of execution is not guaranteed
	// the result of the program is not deterministic
	// or equal for each execution. Everytime they appear
	// in a different order, the result is different.
	go add()
	go print()

	time.Sleep(time.Second)
	fmt.Printf("Final counter: %d\n", counter)
}
