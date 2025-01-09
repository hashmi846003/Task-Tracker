package handlers

import (
	"fmt"
	"strconv"
	"task-tracker/utils"
)

func DeleteTask(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: task-cli delete <id>")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid task ID.")
		return
	}

	tasks, err := utils.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	index := -1
	for i, task := range tasks {
		if task.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Printf("No task found with ID: %d\n", id)
		return
	}

	// Remove the task from the slice
	tasks = append(tasks[:index], tasks[index+1:]...)

	err = utils.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Printf("Task deleted successfully (ID: %d).\n", id)
}
