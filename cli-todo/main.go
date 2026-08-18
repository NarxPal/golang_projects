package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Task struct {
	Id          int
	Task        string
	IsCompleted bool
	CreatedOn   string
}

var taskId int = 0
var taskList []Task

func createTask(taskTitle string) []Task {
	for _, task := range taskList {
		if task.Id > taskId {
			taskId = task.Id
		}
	}

	taskId++ // increase by one
	newTask :=
		Task{
			Id:          taskId,
			Task:        taskTitle,
			IsCompleted: false,
		}
	taskList = append(taskList, newTask)
	return taskList
}

func showAllTask() {
	if len(taskList) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("\nTasks:")
	fmt.Println("────────────────────────────────────────────────────────────")

	for _, task := range taskList {
		icon := "☐"
		status := "incomplete"

		if task.IsCompleted {
			icon = "✓"
			status = "completed"
		}

		fmt.Printf("%s  ID: %-3d  Task: %-30s  Status: %s\n",
			icon,
			task.Id,
			task.Task,
			status,
		)
	}

	fmt.Println("────────────────────────────────────────────────────────────")
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

func toggleIsCompleted(taskId int) {
	for i, task := range taskList {
		if task.Id == taskId {
			taskList[i].IsCompleted = !taskList[i].IsCompleted
		}
	}

}

func getTaskFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return filepath.Join(home, ".local", "share", "todo", "tasks.json")
}

func saveTasks() {
	path := getTaskFilePath()

	MkdirErr := os.MkdirAll(filepath.Dir(path), 0755)
	if MkdirErr != nil {
		fmt.Println("Error creating directory:", MkdirErr)
		return
	}
	jsonBytes, jsonBytesErr := json.MarshalIndent(taskList, "", "    ")
	if jsonBytesErr != nil {
		log.Fatalf("Error parsing JSON: %v", jsonBytesErr)
	}
	err := os.WriteFile(path, jsonBytes, 0644)

	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

}

func loadTasks() {
	path := getTaskFilePath()

	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	jsonErr := json.Unmarshal(data, &taskList)

	if jsonErr != nil {
		log.Fatalf("Error parsing JSON: %v", jsonErr)
	}

}

func main() {
	loadTasks()
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
			saveTasks()

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
				saveTasks() // save to disk file
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
				saveTasks() // save to disk file
			}
		} else {
			// show status of task, by asking user's which task status they want

			fmt.Println("enter taskId to toggle isCompleted:")

			if !scanner.Scan() {
				break
			}

			taskId := strings.TrimSpace(scanner.Text())
			id, err := strconv.Atoi(taskId)

			if err != nil {
				fmt.Printf("no such id- %v present\n", id)
			} else {
				toggleIsCompleted(id)
				saveTasks() // save to disk file
			}
		}

	}

}
