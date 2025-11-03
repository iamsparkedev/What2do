package main

import (
	"fmt"
	"os"
)

// CLI TO-DO APP Without any external libraries
func main() {

    defer fmt.Println("                     ")
    defer fmt.Println("=========================================================")
	defer fmt.Println("4. todo exit       - Exit from the application")
	defer fmt.Println("3. todo delete <taskid>      - Delete a task")
	defer fmt.Println("2. todo list             - View all tasks")
	defer fmt.Println(`1. todo add "<task>"       - Add a new task`)
	defer fmt.Println("Available Commands:")
	defer fmt.Println("                     ")
	defer fmt.Println("Let's get started!")
	defer fmt.Println("You can add, view, and delete your tasks easily.")
	defer fmt.Println("Welcome to What-2-Do, your simple CLI to-do application!")
	defer fmt.Println("======================== WHAT-2-DO ======================")
	defer fmt.Println("                     ")

	if len(os.Args) < 2 {
		fmt.Println("Error: No command provided")
		return
	}

	cmd := os.Args[1]
	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task to add")
			return
		}
		handleAddTask(os.Args[2])
	case "list":
		handleListTasks()
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task ID to delete")
			return
		}
		handleDeleteTask(os.Args[2])
	case "exit":
		fmt.Println("Goodbye!")
		return
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		fmt.Println("Available commands: add, list, delete, exit")
	}
}

