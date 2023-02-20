package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// write a program that props the user to enter integers and store them in an ordered slice
	// the program should be written as a loop
	var numbers []int = make([]int, 0, 3)

	// loop over input
	for {
		fmt.Println("Enter an integer or letter 'x' to exit: ")
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
		if strings.ToLower(input) == "x" {
			fmt.Println("Exiting...")
			break
		}
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Not a valid input, try again")
			continue
		}
		numbers = append(numbers, number)
		sort.Ints(numbers)
		// fmt.Println("Numbers: ", numbers)
	}
}
