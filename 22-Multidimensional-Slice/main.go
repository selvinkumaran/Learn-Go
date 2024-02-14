package main

import "fmt"

func main() {

	ks := []string{"karan", "kumar", "Male", "cricket"}
	ds := []string{"dinesh", "kumar", "Male", "volleyball"}
	ss := []string{"sunil", "kumar", "Male", "kabaddi"}
	ms := []string{"mani", "kumar", "Male", "football"}

	fmt.Println(ks)
	fmt.Println(ds)
	fmt.Println(ss)
	fmt.Println(ms)

	xxs := [][]string{ks, ds, ss, ms}

	fmt.Println(xxs)
}
