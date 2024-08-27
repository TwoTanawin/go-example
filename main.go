package main

import (
	"fmt"
)

// struc
type Student struct {
	Name   string
	Weight int
	Height int
	Grade  string
}

// obj method
func (s Student) Info() string {
	return s.Name + ", grade : " + s.Grade
}

// interface
type Speaker interface {
	Speake() string
}

type Dog struct {
	Name string
}

func (d Dog) Speake() string {
	return "Woof!"
}

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

	//Dynamic array
	mySlice := []int{10, 20, 30, 40, 50}
	fmt.Printf("Cappacity : %d\n", cap(mySlice))
	for i := 0; i <= len(mySlice)-1; i++ {
		// mySlice = append(mySlice, 4)
		fmt.Printf("My Slice : %d\n", mySlice[i])
	}

	mySlice = append(mySlice, 6)
	for i := 0; i <= len(mySlice)-1; i++ {
		fmt.Printf("My Slice : %d\n", mySlice[i])
	}

	//Map (Dict)
	myMap := make(map[string]int)

	myMap["apple"] = 5
	myMap["banana"] = 6
	myMap["orange"] = 7

	fmt.Println("myMap: ", myMap)

	delete(myMap, "banana")

	for key, value := range myMap {
		fmt.Printf("%s -> %d\n", key, value)
	}

	val, ok := myMap["pear"]
	if ok {
		fmt.Printf("Pear is value: %d\n", val)
	} else {
		fmt.Printf("Pear not found in Map\n")
	}

	var student1 [3]Student
	student1[0].Name = "Two ISE"
	student1[0].Height = 172
	student1[0].Weight = 70
	student1[0].Grade = "A"
	fmt.Printf("Struden : %s %d %d %s\n", student1[0].Name, student1[0].Height, student1[0].Weight, student1[0].Grade)

	studentInfo := student1[0].Info()
	fmt.Printf("Stdent info : %s ", studentInfo)
}
