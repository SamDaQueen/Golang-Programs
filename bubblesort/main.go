package main

import (
	sort "bubblesort/bubblesort"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()	{

	// take input of up to 10 integers from the user
	sli := make([]int, 0, 10)
	var input string
	fmt.Println("Please enter up to 10 integers separated by space: ")
	
	// create a scanner and read the line of text
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	// convert string of space-separated integers to slice of integers
	fields := strings.Fields(input)
	for _, num := range fields {
		num, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Please enter space-separated integers!")
			continue
		}
		sli =append(sli, num)
		if len(sli)	>= 10	{
			break
		}
	}

	sort.BubbleSort(sli)
	fmt.Print(sli)
}