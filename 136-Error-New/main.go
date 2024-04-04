// Package Error
package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

var errNorgateMath = errors.New("norgate math: square root of negative value")

func main() {

	fmt.Printf("Type of errNorgateMath %T\n", errNorgateMath)

	_, err := squareRoot(-4)

	if err != nil {
		log.Fatalln(err)
	}

}

// SquareRoot Function to square a given number
func squareRoot(number float64) (float64, error) {

	if number < 0 {
		return 0, errNorgateMath
	}
	return math.Pow(number, 2), nil

}
