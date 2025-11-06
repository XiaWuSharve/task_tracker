package tasktracker

import (
	"time"
)

type Task struct {
	Id          int       `json:"id"`          //A unique identifier for the task
	Description string    `json:"description"` //A short description of the task
	Status      Status    `json:"status"`      //The status of the task (todo, in-progress, done)
	CreatedAt   time.Time `json:"created_at"`  //The date and time when the task was created
	UpdatedAt   time.Time `json:"updated_at"`  //The date and time when the task was last updated
}

type TaskRepository []Task

type TaskRepositoryAction interface {
	Add(task Task)
	Update(source Task, target Task)
	Delete(task Task)
	MarkAs(task Task, status Status)
	ListAll() []Task
	ListBy(status Status) []Task
}

func (t *TaskRepository) Add(task Task) *TaskRepository {
	*t = append(*t, task)
	return t
}
