package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type Person struct {
		fname string
		lname string
	}

	var people []Person

	wd, _ := os.Getwd()
	fmt.Printf("Enter the name for the file you want to read with ext (test.txt): ")

	var fileName string
	_, err := fmt.Scan(&fileName)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	file, err := os.Open(fmt.Sprintf("%s/%s", wd, fileName))
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		err := scanner.Err()
		if err != nil {
			fmt.Println("Error reading name in file:", err)
			break
		}

		line := scanner.Text()
		names := strings.Split(line, " ")

		people = append(people, Person{fname: names[0], lname: names[1]})
	}
	for _, person := range people {
		fmt.Printf("first name: %s\nlast name: %s\n----------------------\n", person.fname, person.lname)
	}
}
