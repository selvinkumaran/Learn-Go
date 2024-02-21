package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("selvin")
	fmt.Fprintln(os.Stdout, "selvin")
	io.WriteString(os.Stdout, "selvin")
}
