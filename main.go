package main

import (
	"fmt"
)

func main() {
	fullName := "TwoISE"
	age := 24
	fmt.Printf("Hello %s age: %d", fullName, age)

	fullName = "Nice Two"

	age *= age

	fmt.Println(fullName, age)

	//for loop
	fmt.Printf("for loop\n")
	for i := 1; i <= 10; i++ {
		fmt.Printf("number : %d\n", i)
	}

	// do while
	fmt.Printf("do while loop\n")
	j := 0
	for {
		fmt.Printf("number : %d\n", j)
		j++
		if j > 10 {
			break
		}
	}

	//while loop
	fmt.Printf("while loop\n")
	k := 1
	for k <= 10 {
		fmt.Printf("number : %d\n", k)
		k++
	}
}
