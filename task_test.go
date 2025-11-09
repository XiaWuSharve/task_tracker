package main

import (
	"fmt"
	"testing"
	"time"
)

func TestAddTask(t *testing.T) {
	var tasks TaskRepository
	tasks.AddTask(Task{
		Id:          1,
		Description: "hello world",
		Status:      IN_PROGRESS,
		CreatedAt:   JSONTimeNow(),
		UpdatedAt:   JSONTimeNow(),
	})
	expected := 1
	if len(tasks) != expected {
		t.Errorf("expected: %d, got len of tasks: %d", expected, len(tasks))
	}
}

func TestUpdateTask(t *testing.T) {
	var tasks TaskRepository
	tasks.AddTask(Task{
		Id:          1,
		Description: "hello world",
		Status:      IN_PROGRESS,
		CreatedAt:   JSONTimeNow(),
		UpdatedAt:   JSONTimeNow(),
	})

	expected := Task{
		Id:          1,
		Description: "hello Sharve",
		Status:      DONE,
		CreatedAt:   JSONTime(time.Date(2222, 12, 12, 12, 12, 12, 12, time.Now().Location())),
		UpdatedAt:   JSONTime(time.Date(2222, 12, 12, 12, 12, 12, 12, time.Now().Location())),
	}

	tasks.Update(1, Task{
		Id:          1,
		Description: "hello Sharve",
		Status:      DONE,
		CreatedAt:   JSONTime(time.Date(2222, 12, 12, 12, 12, 12, 12, time.Now().Location())),
		UpdatedAt:   JSONTime(time.Date(2222, 12, 12, 12, 12, 12, 12, time.Now().Location())),
	})

	if len(tasks) != 1 {
		t.Errorf("expected: %d, got len of tasks: %d", 1, len(tasks))
	}

	if tasks[0] != expected {
		t.Errorf("expected: %s, got %s", fmt.Sprint(expected), fmt.Sprint(tasks[0]))
	}
}

func TestDeleteTask(t *testing.T) {
	var tasks TaskRepository
	tasks.AddTask(Task{
		Id:          1,
		Description: "hello world",
		Status:      IN_PROGRESS,
		CreatedAt:   JSONTimeNow(),
		UpdatedAt:   JSONTimeNow(),
	})

	tasks.AddTask(Task{
		Id:          2,
		Description: "hello Sharve",
		Status:      TODO,
		CreatedAt:   JSONTimeNow(),
		UpdatedAt:   JSONTimeNow(),
	})

	targetTasks := Task{
		Id:          2,
		Description: "hello Sharve",
		Status:      TODO,
		CreatedAt:   JSONTimeNow(),
		UpdatedAt:   JSONTimeNow(),
	}

	tasks.Delete(1)

	if len(tasks) != 1 {
		t.Errorf("expected: %d, got len of tasks: %d", 1, len(tasks))
	}

	if tasks[0] != targetTasks {
		t.Errorf("expected: %s, got %s", fmt.Sprint(targetTasks), fmt.Sprint(tasks[0]))
	}

	tasks.Delete(2)

	if len(tasks) != 0 {
		t.Errorf("expected: %d, got len of tasks: %d", 0, len(tasks))
	}
}

func TestMarkAsTask(t *testing.T) {
	var tasks TaskRepository
	tasks.AddTask(Task{
		Id:          1,
		Description: "hello world",
		Status:      IN_PROGRESS,
		CreatedAt:   JSONTimeNow(),
		UpdatedAt:   JSONTimeNow(),
	})

	tasks.MarkAs(1, DONE)

	expected := Task{
		Id:          1,
		Description: "hello world",
		Status:      DONE,
		CreatedAt:   JSONTimeNow(),
		UpdatedAt:   JSONTimeNow(),
	}

	if tasks[0].Status != expected.Status {
		t.Errorf("expected: %s, got %s", fmt.Sprint(expected.Status), fmt.Sprint(tasks[0]))
	}
}

func TestStoreTasks(t *testing.T) {
	var tasks TaskRepository
	tasks.AddTask(Task{
		Id:          1,
		Description: "hello world",
		Status:      IN_PROGRESS,
		CreatedAt:   JSONTimeNow(),
		UpdatedAt:   JSONTimeNow(),
	})

	tasks.AddTask(Task{
		Id:          2,
		Description: "hello Sharve",
		Status:      TODO,
		CreatedAt:   JSONTimeNow(),
		UpdatedAt:   JSONTimeNow(),
	})

	SaveTasksToFile(tasks, "storage.json")

	loadedTasks, err := LoadTasksFromFile("storage.json")
	if err != nil {
		panic(err)
	}

	loadedTime := time.Time(loadedTasks[0].CreatedAt).Format(time.DateTime)
	expectedTime := time.Time(tasks[0].CreatedAt).Format(time.DateTime)

	if loadedTime != expectedTime {
		t.Errorf("loadedTime: \n%v\n expectedTime: \n%v", loadedTasks, tasks)
	}
}

func TestRune(t *testing.T) {
	t.Errorf("len: %d", len([]rune("⚙️ ")))
}
