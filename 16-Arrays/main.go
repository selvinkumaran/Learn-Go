
package main

import "fmt"

func main() {

	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	name := [...]string{"selvin", "sel", "vin", "cause"}
	fmt.Println(name)

	var new [6]int

	new[1] = 34

	new[4] = 45
	fmt.Println(new)

	exe:= [...]string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)",
	 "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)",
	 "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie",
	"Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)",
	"Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "MadagascarVanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", 
	"Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough",
	 "RaspberrySorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", 
	 "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}
	fmt.Println(exe)
	fmt.Printf("the length of he exe is %v \t The type is %T",len(exe),exe)
}
