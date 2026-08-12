package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// remove extra logs(such as date, time) provided by log package and rather use greetings as prefix
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println("Hello, world!") // fmt.Println, prints to stdout

	message, err := greetings.Hellos([]string{"arora", "soumya", "jay", "akshika"})
	// if error returned than print it to console and exit program
	if err != nil {
		log.Fatal(err)
	}

	// else if no error was returned than print msg to console
	fmt.Println(message)

}
