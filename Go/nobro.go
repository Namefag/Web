package main

import (
	"fmt"
)

func poop() {
	var x = 5
	var y = 7
	if x > y {
		fmt.Println("Yes!")
	} else if x < y {
		fmt.Println("Nigger.")
		fmt.Println("You eat rocks.")
	} else if x == y {
		fmt.Println("No.")
	}
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Cool!")
		fmt.Println("It works.")
		poop()
	}
}
