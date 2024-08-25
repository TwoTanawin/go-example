package main

import (
	"fmt"

	twoISE "github.com/TwoTanawin/go-example.git/twoise"
	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Printf("Hello Go 2")
	fmt.Printf("UUID: %s", id)
	twoISE.SeyHelloTwoISE()
	twoISE.SeyTest()
}
