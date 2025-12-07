package main

import (
	"fmt"
	"os"
)

func handleAddTask(task string) {
	defer fmt.Printf("Task added: %s\n", task)
	

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
