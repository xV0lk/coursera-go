package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	im := map[string]string{
		"name":    "",
		"address": "",
	}
	for k := range im {
		fmt.Printf("Enter your %s: ", k)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
		im[k] = input
	}
	jsonObj, err := json.Marshal(im)
	if err != nil {
		fmt.Println("An error occurred while creating a json. Please try again", err)
	}
	fmt.Println(string(jsonObj))
}
