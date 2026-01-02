package main

import (
	"fmt"
)

func Helpcmd() {
    banner := "_  _  _  _   __  ____     ____      ____   __  \n" +
"╱ )( ╲╱ )( ╲ ╱ _╲(_  _)___(___ ╲ ___(    ╲ ╱  ╲ \n" +
"╲ ╱╲ ╱) __ (╱    ╲ )( (___)╱ __╱(___)) D ((  O )\n" +
"(_╱╲_)╲_)(_╱╲_╱╲_╱(__)    (____)    (____╱ ╲__╱\n"


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
	defer fmt.Println("=========================================================")
	defer fmt.Println(banner)
	defer fmt.Println("                     ")

}