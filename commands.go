package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Task struct
type Task struct {
	Title       string
	Description string
	Completed   bool 
}
// error handler
func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
// file name generator
func generateFileName() string {
	return time.Now().Format("Mon-Jan-2-2006-15_04_05")
}
// Function for adding task.
func addTask(t Task) {
	Helpcmd()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter task Title: ")
	t.Title, _ = reader.ReadString('\n')
	t.Title = strings.TrimSpace(t.Title)
	fmt.Print("Enter task Description: ")
	t.Description, _ = reader.ReadString('\n')
	t.Description = strings.TrimSpace(t.Description)
	fmt.Println("Adding task:", t.Title)
	fmt.Println("Generated filename:", generateFileName())
	file, err := os.Create("./tmp/" + generateFileName() + ".txt")
	errorCheck(err)
	write, err := file.WriteString(generateFileName() + "\n" + "Task: " + t.Title + "\n" + "Description: " + t.Description + "\n")
	errorCheck(err)
	fmt.Printf("Task added: [%d bytes written]\n", write)
	defer file.Close()
}

// Placeholder function for listing tasks
func handleListTasks() {
	Helpcmd()
	readdir, err := os.ReadDir("./tmp/")
	errorCheck(err)
	if len(readdir) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	fmt.Println("You have", len(readdir), "tasks.")
	fmt.Println("=========================================================")
	for _, entry := range readdir {
        if entry.IsDir() {
            continue
        }

        data, err := os.ReadFile("./tmp/" + entry.Name())
        errorCheck(err)

        fmt.Println(string(data))
        fmt.Println("------------------------------")
    }

    fmt.Println("=========================================================")


}

func handleDeleteTask(s string) {
	Helpcmd()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Delete All Tasks ? : 1" + "\n" + "Delete Specific Task ? : 2" + "\n" + "Enter choice: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "1" {
		fmt.Println("Deleting all tasks...")
		readdir, err := os.ReadDir("./tmp/")
		if err != nil {
			fmt.Println("Error reading tasks directory:", err)
			return
		}
		for _, entry := range readdir {
			if entry.IsDir() {
				continue
			}
			err := os.Remove("./tmp/" + entry.Name())
			errorCheck(err)
		}
		fmt.Println("All tasks deleted.")
	} else {
		fmt.Print("Enter the filename of the task to delete: ", s)
		filename, _ := reader.ReadString('\n')
		filename = strings.TrimSpace(filename)
		err := os.Remove("./tmp/" + filename)
		if err != nil {
			fmt.Println("Can't find specified task"+"\n"+"Error deleting task:", err)
			return
		}
		fmt.Println("Task", filename, "deleted.")
	}
}