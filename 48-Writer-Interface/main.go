package main

import (
	"log"
	"os"
)

func main() {

	f, err := os.Create("sample.txt")

	if err != nil {
		log.Fatalf("error %s", err)
	}

	defer f.Close()

	s := []byte("hello selvin")

	_, err = f.Write(s)

	if err != nil {
		log.Fatalf("error %s", err)
	}

}
