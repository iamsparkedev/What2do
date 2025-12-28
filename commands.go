package main

import (
	"fmt"
	"os"
)

func handleAddTask(task string) {
	defer fmt.Printf("Task added: %s\n", task)
    filename := "task.txt"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()
	os.WriteFile(filename, []byte(task), 0644)
}

func handleListTasks() {
	fmt.Println("Listing all tasks...waiting for implementation")
	file,_ := os.Open("tasks.json")
	defer file.Close()
	read,_ := os.ReadFile("tasks.json")
	fmt.Println(string(read))
}

func handleDeleteTask(taskID string) {
	fmt.Printf("Deleting task with ID: %s\n", taskID)
}
