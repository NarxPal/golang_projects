package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	fmt.Println("Hello, world!")

	message := greetings.Hello("arora")
	fmt.Println(message)

}
