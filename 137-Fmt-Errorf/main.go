package main

import (
	"fmt"
	"log"
	"math"
)

var errNorgateMath = fmt.Errorf("norgate math: square root of negative value")

func main() {

	fmt.Printf("Type of errNorgateMath %T\n", errNorgateMath)

	_, err := squareRoot(-4)

	if err != nil {
		log.Fatalln(err)
	}

}

func squareRoot(number float64) (float64, error) {

	if number < 0 {
		return 0, errNorgateMath
	}
	return math.Pow(number, 2), nil

}
