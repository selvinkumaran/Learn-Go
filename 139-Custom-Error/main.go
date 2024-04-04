// Package main provides a simple program to demonstrate JSON marshaling in Go.
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// person struct represents a person with their first name, last name, and sayings.
type person struct {
	First   string   // First name of the person
	Last    string   // Last name of the person
	Sayings []string // Slice of strings containing sayings associated with the person
}

// main function creates an instance of person struct, marshals it into JSON format,
// and prints the JSON representation to the standard output.
func main() {
	// Creating an instance of person struct.
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	// Marshaling the person struct to JSON.
	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)
	}

	// Printing the JSON representation.
	fmt.Println(string(bs))
}

// toJSON function marshals any value into JSON format.
func toJSON(a interface{}) ([]byte, error) {
	// Marshal the given value to JSON format.
	bs, err := json.Marshal(a)

	// Return error if marshaling fails.
	if err != nil {
		return []byte{}, fmt.Errorf("there was an error in toJSON: %v", err)
	}

	// Return the marshaled JSON bytes.
	return bs, nil
}
