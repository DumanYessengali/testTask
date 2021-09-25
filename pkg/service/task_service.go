package service

import (
	"errors"
	_ "github.com/lib/pq"
	tasks "newProjectFolders/testTask1"
	"newProjectFolders/testTask1/pkg/repository"
)

type TaskService struct {
	repository repository.Task
}

func NewTaskService(repository repository.Task) *TaskService {
	return &TaskService{repository: repository}
}

func (t *TaskService) CreateTask(task tasks.Tasks) (*tasks.Tasks, error) {
	return t.repository.CreateTask(task)
}

func (t *TaskService) GetTaskById(id int) (tasks.Tasks, error) {
	return t.repository.GetTaskById(id)
}

func (t *TaskService) UpdateTaskById(id int, name, description, author *string, task tasks.Tasks) error {
	didUpdate := false
	if name != nil {
		task.Name = *name
		didUpdate = true
	}

	if description != nil {
		task.Description = *description
		didUpdate = true
	}

	if author != nil {
		task.Author = *author
		didUpdate = true
	}

	if !didUpdate {
		return errors.New("no update done")
	}

	return t.repository.UpdateTaskById(id, task)
}

func (t *TaskService) DeleteTaskById(id int) error {
	return t.repository.DeleteTaskById(id)
}
