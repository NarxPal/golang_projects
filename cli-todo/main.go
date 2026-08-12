package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	Id          int
	Title       string
	Status      string
	IsCompleted bool
	CreatedOn   string
}

func main() {
	for {

		var userInput string

		fmt.Println("type what u want to do for task: add, edit, status, delete, show")

		_, err := fmt.Scanln(&userInput)
		scanner := bufio.NewScanner(os.Stdin)

		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// if userInput doesn't match the suggested input than let the user re-type
		if userInput != "add" && userInput != "edit" && userInput != "status" && userInput != "delete" && userInput != "show" {
			fmt.Println("Invalid command!")
			continue
		}
		fmt.Println("add task:")

		if !scanner.Scan() {
			break
		}

		task := scanner.Text()

		if task == "" {
			fmt.Println("You didn't type anything, plz try again")
			continue
		}

		fmt.Printf("task added: %v\n", task)

		break
	}

}
