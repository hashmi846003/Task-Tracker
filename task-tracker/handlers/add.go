package handlers

import (
	"fmt"
	"strings"
	"time"
	"task-tracker/models"
	"task-tracker/utils"
)

func AddTask(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: task-cli add <description>")
		return
	}

	description := strings.Join(args, " ")
	tasks, err := utils.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	task := models.Task{
		ID:          utils.GenerateID(),
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, task)
	err = utils.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving task:", err)
		return
	}

	fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
}
