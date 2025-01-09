package main

import (
	"fmt"
	"os"
	"strings"
	"task-tracker/handlers"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch strings.ToLower(command) {
	case "add":
		handlers.AddTask(args)
	case "update":
		handlers.UpdateTask(args)
	case "delete":
		handlers.DeleteTask(args)
	case "mark-in-progress":
		handlers.MarkInProgress(args)
	case "mark-done":
		handlers.MarkDone(args)
	case "list":
		handlers.ListTasks(args)
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}
