package main

import (
	"encoding/json"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func SaveTasksToFile(tasks TaskRepository, filename string) error {
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, jsonData, 0644)
}

func LoadTasksFromFile(filename string) (TaskRepository, error) {

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if os.IsNotExist(err) {
		err = SaveTasksToFile(make(TaskRepository, 0), filename)
		checkError(err)
	} else {
		checkError(err)
		err = file.Close()
		checkError(err)
	}

	data, err := os.ReadFile(filename)
	checkError(err)

	tasks := make(TaskRepository, 0)
	err = json.Unmarshal(data, &tasks)
	checkError(err)

	return tasks, err
}
