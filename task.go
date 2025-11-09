package main

import (
	"errors"
	"fmt"
	"time"
)

type Task struct {
	Id          int      `json:"id"`          //A unique identifier for the task
	Description string   `json:"description"` //A short description of the task
	Status      Status   `json:"status"`      //The status of the task (todo, in-progress, done)
	CreatedAt   JSONTime `json:"created_at"`  //The date and time when the task was created
	UpdatedAt   JSONTime `json:"updated_at"`  //The date and time when the task was last updated
}

type TaskRepository []Task

type TaskRepositoryAction interface {
	AddTask(task Task) int
	AddFromDescription(desc string) int
	Update(sourceId int, target Task) error
	UpdateFromDescription(sourceId int, desc string) error
	Delete(taskId int) error
	MarkAs(taskId int, status Status) error
	ListAll() TaskRepository
	listBy(status Status) TaskRepository
	ListDone() TaskRepository
	ListNotDone() TaskRepository
	ListInProgress() TaskRepository
	findOneById(id int) int
	genUniqueId() int
}

func (t *TaskRepository) genUniqueId() int {
	maxId := 0
	for _, task := range *t {
		maxId = max(maxId, task.Id)
	}
	return maxId + 1
}

func (t *TaskRepository) findOneById(id int) int {
	for i := range *t {
		if (*t)[i].Id == id {
			return i
		}
	}
	return -1
}

func (t *TaskRepository) AddTask(task Task) int {
	task.Id = t.genUniqueId()
	task.Status = TODO
	task.CreatedAt = JSONTime(time.Now())
	task.UpdatedAt = JSONTime(time.Now())
	*t = append(*t, task)
	return task.Id
}

func (t *TaskRepository) AddFromDescription(desc string) int {
	task := Task{
		Description: desc,
	}
	return t.AddTask(task)
}

func (t *TaskRepository) Update(sourceId int, target Task) error {
	return errors.New("update: deprecated")
	existingTaskIndex := t.findOneById(sourceId)
	if existingTaskIndex == -1 {
		errMsg := fmt.Sprintf("Except existing task %d, found no existing task", sourceId)
		return errors.New(errMsg)
	}

	target.UpdatedAt = JSONTime(time.Now())
	(*t)[existingTaskIndex] = target

	return nil
}

func (t *TaskRepository) UpdateFromDescription(sourceId int, desc string) error {
	existingTaskIndex := t.findOneById(sourceId)
	if existingTaskIndex == -1 {
		errMsg := fmt.Sprintf("Except existing task %d, found no existing task", sourceId)
		return errors.New(errMsg)
	}

	(*t)[existingTaskIndex].Description = desc
	(*t)[existingTaskIndex].UpdatedAt = JSONTime(time.Now())

	return nil
}

func (t *TaskRepository) Delete(taskId int) error {
	existingTaskIndex := t.findOneById(taskId)
	if existingTaskIndex == -1 {
		errMsg := fmt.Sprintf("Except existing task %d, found no existing task", taskId)
		return errors.New(errMsg)
	}

	*t = append((*t)[:existingTaskIndex], (*t)[existingTaskIndex+1:]...)

	return nil
}

func (t *TaskRepository) MarkAs(taskId int, status Status) error {
	existingTaskIndex := t.findOneById(taskId)
	if existingTaskIndex == -1 {
		errMsg := fmt.Sprintf("Except existing task %d, found no existing task", taskId)
		return errors.New(errMsg)
	}

	(*t)[existingTaskIndex].UpdatedAt = JSONTime(time.Now())
	(*t)[existingTaskIndex].Status = status

	return nil
}

func (t *TaskRepository) ListAll() TaskRepository {
	dest := make(TaskRepository, len(*t))
	copy(dest, *t)
	return dest
}

func (t *TaskRepository) listBy(status Status) TaskRepository {
	dest := make(TaskRepository, 0, len(*t))
	for _, task := range *t {
		if task.Status == status {
			dest = append(dest, task)
		}
	}
	return dest
}

func (t *TaskRepository) ListDone() TaskRepository {
	return t.listBy(DONE)
}

func (t *TaskRepository) ListNotDone() TaskRepository {
	return append(t.listBy(TODO), t.listBy(IN_PROGRESS)...)
}
