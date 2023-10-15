package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main()	{

	var input string

	// assuming empty slice with with no elements (len 0) and capacity 3
	sli := make([]int, 0, 3)

	for {

		fmt.Printf("Enter integer to be added to slice, or x to exit: ")
		fmt.Scan(&input)

		// check for x to break out of loop
		if input == "X" || input == "x" {
			break
		}

		// convert string input to int if possible
		num, err := strconv.Atoi(input)
		if err != nil	{
			fmt.Printf("Please enter an integer or x!")
			continue
		}

		// add to slice, sort, and display
		sli = append(sli, num)
		sort.Ints(sli)
		fmt.Printf("Sorted slice: %v\n", sli)

	}
}
