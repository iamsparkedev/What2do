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
		addTask(Task{})
		generateFileName()
	case "list":
		handleListTasks()
	case "delete":
		handleDeleteTask("")
	case "exit":
		fmt.Println("Goodbye!")
		return
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		fmt.Println("Available commands: add, list, delete, exit")
		
	}
}