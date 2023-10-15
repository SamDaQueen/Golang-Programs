package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// define a struct for name
	type Name struct {
		fname string
		lname string
	}

	// get the file name from the user
	var fileName string
	fmt.Printf("Please enter the file name: ")
	fmt.Scan(&fileName)

	// creating empty slice for the structs
	sli := make([]Name, 0)

	// open the file and use a scanner to read it line by line
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Print("Error opening file!", err)
		return
	}

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// separate each line into fname and lname, limit to 20, add to struct
		name := strings.Split(line, " ")
		fname, lname := name[0], name[1]
		if len(fname) > 20 {
			fname = fname[:20]
		}
		if len(lname) > 20 {
			lname = lname[:20]
		}
		sli = append(sli, Name{fname, lname})
	}

	file.Close()

	for _, name := range sli {
		fmt.Printf("First name: %s, Last name: %s \n", name.fname, name.lname)
	}

}