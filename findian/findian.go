package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()	{
	fmt.Printf("Enter a string: ")

	// Reader for reading the entire line even with spaces
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')

	if (err != nil)	{
		fmt.Printf("Error found! %v", err)
		return
	}

	str = strings.TrimSpace(strings.ToLower(str))

	if (strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str, "a"))	{
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not found!")
	}

}