package handlers

import (
	"fmt"
	"strconv"
	"time"
	"task-tracker/utils"
)

func MarkInProgress(args []string) {
	markTaskStatus(args, "in-progress")
}

func MarkDone(args []string) {
	markTaskStatus(args, "done")
}

func markTaskStatus(args []string, status string) {
	if len(args) < 1 {
		fmt.Printf("Usage: task-cli mark-%s <id>\n", status)
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

	updated := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
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

	fmt.Printf("Task marked as %s successfully (ID: %d).\n", status, id)
}
