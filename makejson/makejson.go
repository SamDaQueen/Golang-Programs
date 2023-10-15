package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// read the name and address and store in variables
	var name string
	var address string

	fmt.Printf("Please enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Please enter your address: ")
	fmt.Scan(&address)

	// create map and add above values to it
	person := make(map[string]string)
	person["name"] = name
	person["address"] = address

	// convert to JSON and print
	personJSON, err := json.Marshal(person)

	if err != nil {
		print("Error in pasrsing JSON!", err)
		return
	}
	fmt.Println("JSON of the map:", string(personJSON))

}