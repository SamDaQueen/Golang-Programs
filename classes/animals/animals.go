package animals

import "fmt"

type Animal struct {
	food string
	locomotion string
	noise string
}

func NewAnimal(food, locomotion, noise string) *Animal {
	return &Animal{food, locomotion, noise}
}

func (a Animal) Eat() string	{
	fmt.Printf("This animal eats %s\n", a.food)
	return a.food
}

func (a Animal)	Move() string	{
	fmt.Printf("This animal %ss\n", a.locomotion)
	return a.locomotion
}

func (a Animal)	Speak() string	{
	fmt.Printf("This animal makes %s sound\n", a.noise)
	return a.noise
}