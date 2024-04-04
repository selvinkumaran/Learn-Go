package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("name.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	defer f.Close()

	r := strings.NewReader("selvin created a file new")
	io.Copy(f, r)
}
