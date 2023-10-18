package main

import (
	dis "displacement/displacement"
	"fmt"
)

func main() {
	input := GetUserInput()
	fn := dis.GenDisplaceFn(input[0], input[1], input[2])
	fmt.Printf("Displacement for time %.2fs is %.2f", input[3], fn(input[3]))
}

func GetUserInput() [4]float64 {
	var a, v0, s0, t float64
	fmt.Print("Enter acceleration: ")
	fmt.Scan(&a)
	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&v0)
	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&s0)
	fmt.Print("Enter time: ")
	fmt.Scan(&t)

	return [4]float64{a, v0, s0, t}
}