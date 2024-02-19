package main

import (
	"bytes"
	"fmt"
)

func main() {

	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())

	b.WriteString("selvin")
	fmt.Println(b.String())

	b.Reset()
	b.WriteString("New Buffer ")
	fmt.Println(b.String())

	b.Write([]byte("happy happy"))
	fmt.Println(b.String())
}
