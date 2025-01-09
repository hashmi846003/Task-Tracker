package handlers

import (
	"fmt"
	"task-tracker/utils"
)

func ListTasks(args []string) {
	tasks, err := utils.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	filter := ""
	if len(args) > 0 {
		filter = args[0]
	}

	for _, task := range tasks {
		if filter == "" || filter == task.Status {
			fmt.Printf("[%d] %s - %s (Created: %s, Updated: %s)\n", task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	}
}
