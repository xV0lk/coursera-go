package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Printf("-----------------------\nEnter value for acceleration: ")
	var acc float64
	// check if there is an error and ask for a correct value again.
	scanErrorRetry(&acc)
	fmt.Printf("Enter value for initial velocity: ")
	var vel float64
	scanErrorRetry(&vel)
	fmt.Printf("Enter value for initial displacement: ")
	var disp float64
	scanErrorRetry(&disp)
	for n := 0; true; n++ {
		if n == 0 {
			fmt.Printf("-----------------------\nEnter value for time: ")
		}
		if n > 0 {
			fmt.Printf("-----------------------\nEnter another value for time: ")
		}
		var t float64
		scanErrorRetry(&t)
		fn := GenDisplaceFn(acc, vel, disp)
		// displays a count-down from time entered.
		waitCounter(t)
		fmt.Println("Displacement is: ", fn(t))
	}
}

func GenDisplaceFn(acc, vel, disp float64) func(t float64) float64 {

	fn := func(t float64) float64 {
		return (0.5 * acc * math.Pow(t, 2)) + (vel * t) + disp
	}
	return fn
}

// scanErrorRetry accepts a pointer and ask the user for an input.
// If there is an error it will keep asking for a valid input.
func scanErrorRetry(value *float64) {
	_, err := fmt.Scan(value)
	if err != nil {
		fmt.Printf("Please try again with a valid number: ")
		scanErrorRetry(value)
	}
	return
}

// waitCounter accepts a float64 time argument and prints a count-down starting at the value of the time passed.
func waitCounter(t float64) {
	for ts := int(t); ts > 0; ts-- {
		if ts > 1 {
			fmt.Printf("%v...", ts)
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Printf("%v...", ts)
		time.Sleep(1 * time.Second)
		fmt.Printf("\n")
	}
}
