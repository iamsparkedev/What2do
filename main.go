package main

import (
	"fmt"
	"os"
)

// CLI TO-DO APP Without any external libraries
func main() {

	if len(os.Args) == 1 {
		Helpcmd()
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