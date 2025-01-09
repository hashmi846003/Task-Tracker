package utils

import (
	"encoding/json"
	"os"
	"task-tracker/models"
)

const FileName = "tasks.json"

func LoadTasks() ([]models.Task, error) {
	var tasks []models.Task
	file, err := os.OpenFile(FileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, _ := file.Stat()
	if stat.Size() == 0 {
		return tasks, nil
	}

	err = json.NewDecoder(file).Decode(&tasks)
	return tasks, err
}

func SaveTasks(tasks []models.Task) error {
	file, err := os.Create(FileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(tasks)
}
