package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (a Cow) Eat() {
	fmt.Printf("\033[34m%s\033[0m\n\n", a.food)
}

func (a Cow) Move() {
	fmt.Printf("\033[34m%s\033[0m\n\n", a.locomotion)
}

func (a Cow) Speak() {
	fmt.Printf("\033[34m%s\033[0m\n\n", a.noise)
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (a Bird) Eat() {
	fmt.Printf("\033[34m%s\033[0m\n\n", a.food)
}

func (a Bird) Move() {
	fmt.Printf("\033[34m%s\033[0m\n\n", a.locomotion)
}

func (a Bird) Speak() {
	fmt.Printf("\033[34m%s\033[0m\n\n", a.noise)
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (a Snake) Eat() {
	fmt.Printf("\033[34m%s\033[0m\n\n", a.food)
}

func (a Snake) Move() {
	fmt.Printf("\033[34m%s\033[0m\n\n", a.locomotion)
}

func (a Snake) Speak() {
	fmt.Printf("\033[34m%s\033[0m\n\n", a.noise)
}

type animals map[string]Animal

func main() {
	animals := animals{}
	fmt.Printf("\n\033[1;34mRead carefully this instructions!\033[0m\n\n")
	fmt.Println("You have two commands available: \033[4;32mnewanimal\033[0m and \033[4;32mquery\033[0m\n")
	fmt.Println("\033[1mnewanimal:\033[0m Allows you to create a new animal consisting of any name and animal type (cow, bird, snake).")
	fmt.Println("You need to pass 3 arguments separated by a space ex: newanimal jack cow\n")
	fmt.Println("\033[1mquery:\033[0m Allows you to query an existing animal by name and get information about it (eat, move, or speak).")
	fmt.Println("You need to pass 3 arguments separated by a space ex: query jack move\n")
	for {
		fmt.Printf("Please enter a command according to instructions: \n> ")
		opType, anName, meInput, error := scanInput(animals)
		if error != nil {
			printError(error)
			continue
		}
		if opType == "newanimal" {
			newAnimal(meInput, anName, animals)
		}
		if opType == "query" {
			queryAnimal(anName, meInput, animals)
		}
		if error != nil {
			printError(error)
			continue
		}
	}
}

func printError(err error) {
	fmt.Printf("\033[1;31m%s, try again with a valid input.\033[0m\n\n", err)
}

/*
scanInput receives a map as an argument and ask user for an input.

It will verify the information entered by the user is correct.
And will return the validated inputd and an error.
*/
func scanInput(animals animals) (opType, anName, meInput string, sError error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		sError = err
	}
	words := strings.Split(scanner.Text(), " ")
	if len(words) != 3 {
		sError = errors.New("Invalid number of arguments, expected 3")
		return
	}
	opType = strings.ToLower(words[0])
	anName = strings.ToLower(words[1])
	meInput = strings.ToLower(words[2])
	switch opType {
	case "newanimal":
		// check if animal type is valid and is a new one
		if meInput != "cow" && meInput != "snake" && meInput != "bird" {
			sError = errors.New("Not a valid animal, please enter cow, bird or snake as third argument")
			return
		}
		if _, ok := animals[anName]; ok {
			sError = errors.New(fmt.Sprintf("%s already exists, please provide a new name", anName))
			return
		}
	case "query":
		// check if name is already saved
		if _, ok := animals[anName]; !ok {
			sError = errors.New(fmt.Sprintf("No animal named %s created", anName))
			return
		}
		// check if method is valid
		if meInput != "eat" && meInput != "move" && meInput != "speak" {
			sError = errors.New(fmt.Sprintf("%s is not a valid method, please enter eat, move or speak as third argument", meInput))
			return
		}
	default:
		sError = errors.New(fmt.Sprintf("%s is not a valid command, please enter newanimal or query as first argument", opType))
	}
	return opType, anName, meInput, sError
}

func newAnimal(anType, anName string, animals animals) {
	switch anType {
	case "cow":
		animals[anName] = Cow{"grass", "walk", "moo"}
	case "bird":
		animals[anName] = Cow{"worms", "fly", "peep"}
	case "snake":
		animals[anName] = Cow{"mice", "slither", "hsss"}
	}
	fmt.Printf("\033[1;32mCreated it!!\033[0m\n\n")
}

func queryAnimal(anName, method string, animals animals) {
	animal := animals[anName]
	switch method {
	case "eat":
		fmt.Printf("\033[34m%s eats \033[0m", anName)
		animal.Eat()
	case "move":
		fmt.Printf("\033[34m%s can \033[0m", anName)
		animal.Move()
	case "speak":
		fmt.Printf("\033[34m%s speaks \033[0m", anName)
		animal.Speak()
	}
}
