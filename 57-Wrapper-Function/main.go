package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	xb, err := readFile("sample.txt")

	if err != nil {
		log.Fatalf("readFile in main %s", err)
	}

	fmt.Println(xb)
	fmt.Println(string(xb))

}

func readFile(fn string) ([]byte, error) {
	xb, err := os.ReadFile(fn)
	if err != nil {
		return nil, fmt.Errorf("erroe in the readFile func %s", err)
	}
	return xb, nil
}
