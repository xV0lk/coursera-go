package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

type animals map[string]Animal

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.noise
}

func main() {
	animals := animals{
		"cow": {
			food:       "grass",
			locomotion: "walk",
			noise:      "moo",
		},
		"bird": {
			food:       "worms",
			locomotion: "fly",
			noise:      "peep",
		},
		"snake": {
			food:       "mice",
			locomotion: "slither",
			noise:      "hsss",
		},
	}
	// fmt.Printf("\033[1;34mInstructions!\033[0m\nYou have two commands available \033[4;32mnewanimal\033[0m and \033[4;32mquery\033[0m\n\n\033[1mnewanimal:\033[0m Allows you to create a new animal consisting of any name ad animal type (cow, bird, snake).\nYou need to pass 3 arguments separated by a space ex: newanimal jack cow\n\n\033[1mquery:\033[0m Allows you to query by name an existing and give you information about it(eat, move or speak).\nYou need to pass 3 arguments separated by a space ex: query jack move\n\n")
	fmt.Printf("\033[1;34mInstructions!\033[0m\n\n")
	fmt.Println("You have two commands available: \033[4;32mnewanimal\033[0m and \033[4;32mquery\033[0m\n")
	fmt.Println("\033[1mnewanimal:\033[0m Allows you to create a new animal consisting of any name and animal type (cow, bird, snake).")
	fmt.Println("You need to pass 3 arguments separated by a space ex: newanimal jack cow\n")
	fmt.Println("\033[1mquery:\033[0m Allows you to query an existing animal by name and get information about it (eat, move, or speak).")
	fmt.Println("You need to pass 3 arguments separated by a space ex: query jack move\n")

	for {
		fmt.Printf("Please enter an animal and a method: \n> ")
		animal, anInput, method, error := scanInput(animals)
		if error != nil {
			printError(error)
			continue
		}
		error = checkMethod(animal, anInput, method)
		if error != nil {
			printError(error)
			continue
		}
	}
}

func printError(err error) {
	fmt.Printf("%s, please try again with a valid input\n\n", err)
}

/*
scanInput receives a map as an argument and ask user for an input.

It will verify if the animal passed is inside the collection
And will return an instance of the animal, a method and an error.
*/
func scanInput(animals animals) (animal Animal, anInput, meInput string, sError error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		sError = err
	}
	words := strings.Split(scanner.Text(), " ")
	if len(words) != 2 {
		sError = errors.New("Invalid number of arguments, expected only one animal an a method")
		return
	}
	anInput = words[0]
	meInput = words[1]
	animal, ok := animals[anInput]
	if !ok {
		sError = errors.New("Invalid animal, expected cow, bird or snake")
		return
	}
	return animal, anInput, meInput, sError
}

func checkMethod(animal Animal, anInput, method string) (error error) {
	switch strings.ToLower(method) {
	case "eat":
		fmt.Printf("%ss eat %s\n\n", anInput, animal.food)
	case "move":
		fmt.Printf("%ss move by %sing\n\n", anInput, animal.locomotion)
	case "speak":
		fmt.Printf("When %ss speak you will hear '%s'\n\n", anInput, animal.noise)
	default:
		error = errors.New("Invalid method, expected eat, move or speak")
		return
	}
	return
}
