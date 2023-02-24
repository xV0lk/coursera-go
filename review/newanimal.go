package main

import (
	"fmt"
	"strings"
)

const (
	COW   = "cow"
	BIRD  = "bird"
	SNAKE = "snake"
)

// Animals types
var cowInit = Cow{Animals{"grass", "walk", "moo"}}
var birdInit = Bird{Animals{"worms", "fly", "peep"}}
var snakeInit = Snake{Animals{"mice", "slither", "hsss"}}

func main() {
	var requestType, request2, request3 string
	for {
		fmt.Println("Enter request(newanimal or query) animal and action: ")
		fmt.Print(">")
		_, err := fmt.Scanf("%s %s %s", &requestType, &request2, &request3)
		if err != nil {
			fmt.Println("Error typing the request ", err)
			continue
		}
		switch strings.ToLower(requestType) {
		case "newanimal":
			createAnimal(request2, request3)
		case "query":
			query(request2, request3)
		}
	}
}

type Animals struct {
	Food       string
	Locomotion string
	Noise      string
}
type Cow struct {
	Animals
}
type Bird struct {
	Animals
}
type Snake struct {
	Animals
}

func (animal *Animals) Eat() {
	fmt.Println(animal.Food)
}
func (animal *Animals) Move() {
	fmt.Println(animal.Locomotion)
}
func (animal *Animals) Speak() {
	fmt.Println(animal.Noise)
}

func createAnimal(animalName, animalType string) {
	switch strings.ToLower(animalType) {
	case COW:
		setsAnimals[animalName] = cowInit
	case BIRD:
		setsAnimals[animalName] = birdInit
	case SNAKE:
		setsAnimals[animalName] = snakeInit
	}
	fmt.Println("Created it!")
}

func query(animalName, animalAction string) {
	animal := setsAnimals[animalName]
	switch dataAnimal := animal.(type) {
	case Cow:
		switch strings.ToLower(animalAction) {
		case "eat":
			dataAnimal.Eat()
		case "move":
			dataAnimal.Move()
		case "speak":
			dataAnimal.Speak()
		}
	case Bird:
		switch strings.ToLower(animalAction) {
		case "eat":
			dataAnimal.Eat()
		case "move":
			dataAnimal.Move()
		case "speak":
			dataAnimal.Speak()
		}
	case Snake:
		switch strings.ToLower(animalAction) {
		case "eat":
			dataAnimal.Eat()
		case "move":
			dataAnimal.Move()
		case "speak":
			dataAnimal.Speak()
		}
	}

}

type Animal interface {
	Eat()
	Move()
	Speak()
}

var setsAnimals = map[string]interface{}{}
