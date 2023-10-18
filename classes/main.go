package main

import (
	"bufio"
	animals "classes/animals"
	"fmt"
	"os"
	"strings"
)

func main()	{
	fmt.Println("Please type your request after > in the form 'animal info'")
	fmt.Println("For e.g., to know the food a cow eats, type 'cow eat'")
	fmt.Println("Animals: cow, bird, snake. Info: eat, move, speak")
	fmt.Printf("To exit the program at any time, enter x\n\n")

	cow, bird, snake := GenData()

	for {
		fmt.Print(">")
		
		// create a scanner and read the line of text
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		request := strings.Split(input, " ")

		switch(request[0])	{
		case "cow":
			GetResponse(cow, request[1])
		case "bird":
			GetResponse(bird, request[1])
		case "snake":
			GetResponse(snake, request[1])
		default:
			return
		}

	}
}

func GetResponse(a *animals.Animal, req string)	{
	switch(req){
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Println("Invalid request!")
		return
	}
}

func GenData() (*animals.Animal, *animals.Animal, *animals.Animal)	{
	cow := animals.NewAnimal("grass", "walk", "moo")
	bird := animals.NewAnimal("worms", "fly", "peep")
	snake := animals.NewAnimal("mice", "slither", "hsss")

	return cow, bird, snake
}