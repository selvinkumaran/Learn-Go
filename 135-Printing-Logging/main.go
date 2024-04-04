package main

import (
	"os"
)

func main() {
	_, err := os.Open("newfile.txt")
	if err != nil {
		// fmt.Println("Fmt Println Error:", err)
		// log.Println("Log Println Error", err)
		// fmt.Println("HI EVERYONE")
		// log.Fatalln("Log Fatalln Error", err) //the program will be terminate when an error occurs and defer functions also won't be work (Fatal)
		// fmt.Println("HI EVERYONE it won't run")
		// log.Panic("Log panic Error", err)
		panic(err)
	}
}
