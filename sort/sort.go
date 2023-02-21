package main

import "fmt"

func main() {
	var numbers []int = make([]int, 0, 10)
	for count := 1; count <= 10; {
		fmt.Printf("-----------------------\nEnter an integer: ")
		var input int
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		numbers = append(numbers, input)
		count += 1
		bubbleSort(numbers)
		fmt.Println(numbers)
		if count > 10 {
			println("No more space left, program exiting...")
			break
		}
	}
}

// bubbleSort, sorts a slice.
func bubbleSort(s []int) {
	n := len(s)

	var sw bool
	for i := 0; i < n; i++ {
		sw = false
		for j := 0; j < n-i-1; j++ {
			if s[j] > s[j+1] {
				swap(s, j)
				sw = true
			}
		}
		if sw == false {
			break
		}
	}
}

// swap receive as arguments a slice and an index, It will swap the item at the index location with its next sibling
func swap(s []int, i int) {
	s[i], s[i+1] = s[i+1], s[i]
}
