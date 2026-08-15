package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Id          int
	Task        string
	Status      string
	IsCompleted bool
	CreatedOn   string
}

var taskId int = 0
var taskList []Task

func createTask(taskTitle string) []Task {
	taskId++
	newTask :=
		Task{
			Id:          taskId,
			Task:        taskTitle,
			Status:      "in-progress",
			IsCompleted: false,
		}
	taskList = append(taskList, newTask)
	return taskList
}

func showAllTask() {
	fmt.Printf("all tasks %v\n", taskList)
}

func editTaskById(taskId int, scanner *bufio.Scanner) {
	for i, task := range taskList {
		if task.Id == taskId {
			fmt.Println("task title to edit is:")
			fmt.Printf("%v\n", taskList[i].Task)

			if !scanner.Scan() {
				break
			}
			editedTaskTitle := strings.TrimSpace(scanner.Text())
			taskList[i].Task = editedTaskTitle
		}
	}
}

func deleteTaskById(taskId int) {
	for i, task := range taskList {
		if task.Id == taskId {
			taskList = append(taskList[:i], taskList[i+1:]...)
		}
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {

		var userInput string

		fmt.Println(">type what u want to do for task: add, edit, status, delete, show")

		if !scanner.Scan() {
			break
		}
		userInput = strings.TrimSpace(scanner.Text())

		fmt.Print("\033[1A\033[2K") // clear user input after "enter" is pressed

		// if userInput doesn't match the suggested input than let the user re-type
		if userInput != "add" && userInput != "edit" && userInput != "status" && userInput != "delete" && userInput != "show" {
			fmt.Println("Invalid command!")
			continue
		}

		if userInput == "add" {

			fmt.Println("add task:")

			if !scanner.Scan() {
				break
			}

			task := strings.TrimSpace(scanner.Text())

			if task == "" {
				fmt.Println(">You didn't type anything, plz try again")
				continue
			}

			createTask(task)

		} else if userInput == "show" {
			showAllTask()
		} else if userInput == "edit" {
			fmt.Println("enter taskId to edit task:")

			if !scanner.Scan() {
				break
			}

			taskId := strings.TrimSpace(scanner.Text())
			id, err := strconv.Atoi(taskId)

			if err != nil {
				fmt.Printf("no such id- %v present\n", id)
			} else {
				editTaskById(id, scanner)
			}

		} else if userInput == "delete" {
			fmt.Println("enter taskId to delete task:")

			if !scanner.Scan() {
				break
			}
			taskId := strings.TrimSpace(scanner.Text())
			id, err := strconv.Atoi(taskId)

			if err != nil {
				fmt.Printf("no such id- %v present\n", id)
			} else {
				deleteTaskById(id)
			}
		} else {
			// show status of task, by asking user's which task status they want

		}

	}

}
