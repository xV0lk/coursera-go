package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const ROUTINES = 4

func main() {
	var numbers []int
	var endLoop bool
	c := make(chan []int)
	var wg sync.WaitGroup

	for !endLoop {
		fmt.Printf("\033[1;34mEnter an integer or letter 'x' to exit: \033[0m\n> ")
		processInput(&numbers, &endLoop)
	}

	go waitAndClose(c, &wg)
	controller(numbers, c, &wg)
	printer(c)
	fmt.Println("Program end")
}

// ProcessInput scan for inputs from command line and
// process them according to the following rules:
//
// 1. Receives a valid integer and append it to a slice
//
// 2. If letter x is passed, it stops asking for inputs.
//
// 3. If there was an error scanning, it prompts for a valid input
//
// As arguments, it receives a slice to append the values and
// a pointer to a boolean that will represent if the loop needs to end.
func processInput(numbers *[]int, endLoop *bool) {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Printf("\033[1;31mError: \033[;31m%s\n\n\033[1;34mTry again with a valid input: \033[0m\n> ", err)
		processInput(numbers, endLoop)
	}
	// Check if order to stop receiving inputs come
	// and prints the final slice
	if strings.ToLower(input) == "x" {
		*endLoop = true
		fmt.Printf("\033[1;32mStarting sort of: \033[;32m%v\n", *numbers)
		return
	}
	inputNum, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("\033[1;31mError: \033[;31mNot a valid integer\n\n\033[1;34mTry again with a valid input: \033[0m\n> ")
		processInput(numbers, endLoop)
		return
	}
	*numbers = append(*numbers, inputNum)
}

func controller(numbers []int, c chan []int, wg *sync.WaitGroup) {
	var sl []int
	for i := 1; i <= ROUTINES; i++ {
		divideSlice(&sl, numbers, i, ROUTINES)
		wg.Add(1)
		go sortConcurrent(i, sl, c, wg)
	}
}

func sortConcurrent(i int, sl []int, c chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("\033[;34mI'm process #%d and will be sorting: %v\033[0m\n", i, sl)
	sort.Ints(sl)
	fmt.Printf("\033[;34mI'm process #%d and this is the result: %v\033[0m\n", i, sl)
	c <- sl
}

func printer(c chan []int) {
	var newSl []int
	for slice := range c {
		newSl = append(newSl, slice...)
	}
	sort.Ints(newSl)
	fmt.Printf("\033[1;34mWe just finished your sort, here is your result: \033[1;32m%v\033[0m\n", newSl)
}

func waitAndClose(c chan []int, wg *sync.WaitGroup) {
	wg.Wait()
	close(c)
}

func divideSlice(sl *[]int, parent []int, iterator, divisions int) {
	slen := len(parent) / divisions
	switch iterator {
	case 1:
		*sl = parent[0:slen]
	case divisions:
		slStart := (slen * (iterator - 1))
		*sl = parent[slStart:]
	default:
		slStart := (slen * (iterator - 1))
		slEnd := slStart + slen
		*sl = parent[slStart:slEnd]
	}
}
