package main

import "fmt"

func handleAddTask(task string) {
	fmt.Printf("Task added: %s\n", task)
}

func handleListTasks() {
	fmt.Println("Listing all tasks...")
}

func handleDeleteTask(taskID string) {
	fmt.Printf("Deleting task with ID: %s\n", taskID)
}
