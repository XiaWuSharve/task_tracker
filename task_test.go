package tasktracker

import (
	"fmt"
	"testing"
	"time"
)

func TestAddTask(t *testing.T) {
	var tasks TaskRepository
	tasks.Add(Task{
		Id:          1,
		Description: "hello world",
		Status:      inProgress,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	expected := 1
	if len(tasks) != expected {
		t.Errorf("expected: %d, got len of tasks: %d", expected, len(tasks))
	}
}

func TestUpdateTask(t *testing.T) {
	var tasks TaskRepository
	tasks.Add(Task{
		Id:          1,
		Description: "hello world",
		Status:      inProgress,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	expected := Task{
		Id:          1,
		Description: "hello Sharve",
		Status:      done,
		CreatedAt:   time.Date(2222, 12, 12, 12, 12, 12, 12, time.Now().Location()),
		UpdatedAt:   time.Date(2222, 12, 12, 12, 12, 12, 12, time.Now().Location()),
	}

	tasks.Update(1, Task{
		Id:          1,
		Description: "hello Sharve",
		Status:      done,
		CreatedAt:   time.Date(2222, 12, 12, 12, 12, 12, 12, time.Now().Location()),
		UpdatedAt:   time.Date(2222, 12, 12, 12, 12, 12, 12, time.Now().Location()),
	})

	if len(tasks) != 1 {
		t.Errorf("expected: %d, got len of tasks: %d", 1, len(tasks))
	}

	if tasks[0] != expected {
		t.Errorf("expected: %s, got %s", fmt.Sprint(expected), fmt.Sprint(tasks[0]))
	}
}
