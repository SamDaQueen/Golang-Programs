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
	GetName() string
}

type Cow struct {
	name string
}

type Bird struct {
	name string
}

type Snake struct {
	name string
}

func NewAnimal(name, a_type string) Animal {
	switch a_type	{
	case "cow":
		return &Cow{name}
	case "bird":
		return &Bird{name}
	case "snake":
		return &Snake{name}
	default:
		fmt.Println("Animal type can only be cow, bird or snake!")
		return nil
	}
}

var CowData = map[string]string{
	"eat": "grass",
	"move": "walk",
	"speak": "moo",
}

var BirdData = map[string]string{
	"eat": "worms",
	"move": "fly",
	"speak": "peep",
}

var SnakeData = map[string]string{
	"eat": "mice",
	"move": "slither",
	"speak": "hsss",
}

func FindFood(a Animal, request string) string	{
	switch a.(type)	{
	case *Cow:
		val, ok := CowData[request]
		if ok {
			return val
		} else {
			return "Info not found!"
		}
	case *Bird:
		val, ok := BirdData[request]
		if ok {
			return val
		} else {
			return "Info not found!"
		}
	case *Snake:
		val, ok := SnakeData[request]
		if ok {
			return val
		} else {
			return "Info not found!"
		}
	default:
		return "Animal can only be cow bird, or snake!"
	}
}

func (c *Cow) Eat()	{
	if c == nil {
		fmt.Println("No cow exists")
		return
	}
	fmt.Printf("Cow eats grass\n")
}

func (c *Cow) Move()	{
	if c == nil {
		fmt.Println("No cow exists")
	}
	fmt.Printf("Cow walks\n")
}

func (c *Cow) Speak()	{
	if c == nil {
		fmt.Println("No cow exists")
	}
	fmt.Printf("Cow makes moo sound\n")
}

func (c *Cow) GetName() string {
	if c == nil {
		fmt.Println("No cow exists")
	}
	return c.name
}

func (b *Bird) Eat()	{
	if b == nil {
		fmt.Println("No bird exists")
		return
	}
	fmt.Printf("Bird eats worms\n")
}

func (b *Bird) Move()	{
	if b == nil {
		fmt.Println("No bird exists")
	}
	fmt.Printf("Bird flies\n")
}

func (b *Bird) Speak()	{
	if b == nil {
		fmt.Println("No bird exists")
	}
	fmt.Printf("Bird makes peep sound\n")
}


func (b *Bird) GetName() string {
	if b == nil {
		fmt.Println("No bird exists")
	}
	return b.name
}

func (s *Snake) Eat()	{
	if s == nil {
		fmt.Println("No snake exists")
		return
	}
	fmt.Printf("Snake eats mice\n")
}

func (s *Snake) Move()	{
	if s == nil {
		fmt.Println("No snake exists")
	}
	fmt.Printf("Snake slithers\n")
}

func (s *Snake) Speak()	{
	if s == nil {
		fmt.Println("No snake exists")
	}
	fmt.Printf("Snake makes hsss sound\n")
}

func (s *Snake) GetName() string {
	if s == nil {
		fmt.Println("No snake exists")
	}
	return s.name
}

func main()	{
	fmt.Println("Please type your command after >")
	fmt.Println("Command must be 'newanimal' or 'query' type")
	fmt.Printf("To exit the program at any time, enter x\n\n")

	var animals []Animal

	for {
		fmt.Print(">")
		
		// create a scanner and read the line of text
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		request := strings.Split(input, " ")

		switch(request[0])	{
		case "newanimal":
			if len(request) < 3 {
				fmt.Println("The format is 'newanimal <name> <type>")
				continue
			}
			// create a new animal and add to the list
			animals = append(animals, NewAnimal(request[1], request[2]))
			
		case "query":
			if len(request) < 3 {
				fmt.Println("The format is 'query <name> <info>")
				continue
			}
			var found bool = false
			for _, a := range animals {
				name := a.GetName()
				if name == request[1]	{
					fmt.Println(FindFood(a, request[2]))
					found = true
				}
			}
			if !found {
				fmt.Println("Animal not found!")
			}
		default:
			return
		}

	}
}