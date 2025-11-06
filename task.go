package tasktracker

import (
	"errors"
	"fmt"
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
	Add(task Task) (*TaskRepository, error)
	Update(sourceId int, target Task) (*TaskRepository, error)
	Delete(taskId int) (*TaskRepository, error)
	MarkAs(taskId int, status Status) (*TaskRepository, error)
	ListAll() TaskRepository
	listBy(status Status) TaskRepository
	ListDone() TaskRepository
	ListNotDone() TaskRepository
	ListInProgress() TaskRepository
	findOneById(id int) int
}

func (t *TaskRepository) findOneById(id int) int {
	for i := range *t {
		if (*t)[i].Id == id {
			return i
		}
	}
	return -1
}

func (t *TaskRepository) Add(task Task) (*TaskRepository, error) {
	existingTaskIndex := t.findOneById(task.Id)
	if existingTaskIndex != -1 {
		errMsg := fmt.Sprintf("Except no existing task, found task %d on %d", task.Id, existingTaskIndex)
		return t, errors.New(errMsg)
	}
	*t = append(*t, task)
	return t, nil
}

func (t *TaskRepository) Update(sourceId int, target Task) (*TaskRepository, error) {
	existingTaskIndex := t.findOneById(sourceId)
	if existingTaskIndex == -1 {
		errMsg := fmt.Sprintf("Except existing task %d, found no existing task", sourceId)
		return t, errors.New(errMsg)
	}

	(*t)[existingTaskIndex] = target
	return t, nil
}

func (t *TaskRepository) Delete(taskId int) (*TaskRepository, error) {
	existingTaskIndex := t.findOneById(taskId)
	if existingTaskIndex == -1 {
		errMsg := fmt.Sprintf("Except existing task %d, found no existing task", taskId)
		return t, errors.New(errMsg)
	}

	*t = append((*t)[:existingTaskIndex], (*t)[existingTaskIndex+1:]...)

	return t, nil
}

func (t *TaskRepository) MarkAs(taskId int, status Status) (*TaskRepository, error) {
	existingTaskIndex := t.findOneById(taskId)
	if existingTaskIndex == -1 {
		errMsg := fmt.Sprintf("Except existing task %d, found no existing task", taskId)
		return t, errors.New(errMsg)
	}

	(*t)[existingTaskIndex].Status = status

	return t, nil
}

func (t *TaskRepository) ListAll() TaskRepository {
	dest := make(TaskRepository, len(*t))
	copy(dest, *t)
	return dest
}
