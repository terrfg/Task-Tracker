package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli [command]")
		return
	}
	cmmand := os.Args[1]
	tasks, _ := loadTasks()
	switch cmmand {
	case "add":
		desc := os.Args[2]
		id := len(tasks) + 1
		task := Task{ID: id, Description: desc, Status: string(Todo)}
		tasks = append(tasks, task)
		saveTasks(tasks)
		fmt.Println("Task added successfully (ID: %d)", id)
	case "delete":
		id, _ := strconv.Atoi(os.Args[2])
		var newTasks []Task
		for _, t := range tasks {
			if t.ID != id {
				newTasks = append(newTasks, t)
			}
		}
		saveTasks(newTasks)
	case "update":
		id, _ := strconv.Atoi(os.Args[2])
		desc := os.Args[3]
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Description = desc
			}
		}
		saveTasks(tasks)
	case "mark-in-progress":
		id, _ := strconv.Atoi(os.Args[2])
		updateStatus(tasks, id, InProgress)
	case "mark-done":
		id, _ := strconv.Atoi(os.Args[2])
		updateStatus(tasks, id, Done)
	case "list":
		if len(os.Args) == 2 {
			printTask(tasks)
		} else {
			status := Status(os.Args[2])
			printByStatus(tasks, status)
		}
	}
}

func updateStatus(tasks []Task, id int, status Status) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = string(status)
		}
	}
	saveTasks(tasks)
}
func printTask(task []Task) {
	for _, t := range task {
		fmt.Printf("[%d] %s (%s)\n", t.ID, t.Description, t.Status)
	}
}
func printByStatus(tasks []Task, status Status) {
	for i := range tasks {
		if tasks[i].Status == string(status) {
			fmt.Printf("[%d] %s\n", tasks[i].ID, tasks[i].Description)
		}
	}
}
