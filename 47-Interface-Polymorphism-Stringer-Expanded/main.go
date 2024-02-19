package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

type count int

func (b book) String() string {
	return fmt.Sprint("This is a book ", b.title)
}

func (c count) String() string {
	return fmt.Sprint("The number is ", strconv.Itoa(int(c)))

}

func Info(s fmt.Stringer) {
	log.Println("Log from stringer", s.String())
}

func main() {
	b := book{
		"struggle not for struggle",
	}

	var c count = 23

	Info(b)
	Info(c)
}
