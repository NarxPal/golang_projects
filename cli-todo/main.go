package main

import (
	"bufio"
	"encoding/json"
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

var taskId int = 0
var taskList []string

func createTask(taskTitle string) []string {
	taskId++
	newTask :=
		Task{
			Id:          taskId,
			Title:       taskTitle,
			Status:      "in-progress",
			IsCompleted: false,
		}

	jsonBytes, _ := json.Marshal(newTask)
	objectAsString := string(jsonBytes)
	taskList = append(taskList, objectAsString)
	return taskList
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

		taskData := createTask(task)

		fmt.Printf("task added: %v\n", taskData)

		break
	}

}
