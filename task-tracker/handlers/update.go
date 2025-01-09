package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"task-tracker/utils"
)

func UpdateTask(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: task-cli update <id> <new description>")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid task ID.")
		return
	}

	newDescription := strings.Join(args[1:], " ")
	tasks, err := utils.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	updated := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = task.UpdatedAt
			updated = true
			break
		}
	}

	if !updated {
		fmt.Printf("No task found with ID: %d\n", id)
		return
	}

	err = utils.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("Task updated successfully.")
}
