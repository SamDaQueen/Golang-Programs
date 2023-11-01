package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	var third_string, comand, name string
	memory := map[string]string{}

	for {
		fmt.Println(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		words := strings.Fields(scanner.Text())
		comand = strings.ToLower(words[0])
		name = strings.ToLower(words[1])
		third_string = strings.ToLower(words[2])
		if len(words) != 3 {
			fmt.Println("Invalid Input")
			continue
		}

		if comand != "newanimal" && comand != "query" {

			fmt.Println("Invalid Comand")
			continue
		}

		if comand == "newanimal" && third_string != "cow" && third_string != "snake" && third_string != "bird" {

			fmt.Println("Invalid Animal")
			continue
		}

		if comand == "query" && third_string != "eat" && third_string != "move" && third_string != "speak" {

			fmt.Println("Invalid Action")
			continue
		}
		if comand == "newanimal" {
			memory[name] = third_string
			fmt.Println("Created it!")
		} else if comand == "query" {
			var a Animal
			if memory[name] == "cow" {
				a = Cow{}
			} else if memory[name] == "bird" {
				a = Bird{}
			} else if memory[name] == "snake" {
				a = Snake{}
			}
			if third_string == "eat" {
				a.Eat()
			} else if third_string == "move" {
				a.Move()
			} else {
				a.Speak()
			}
		}

	}

}
