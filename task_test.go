package tasktracker

import (
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
