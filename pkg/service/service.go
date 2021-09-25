package service

import (
	_ "github.com/lib/pq"
	tasks "newProjectFolders/testTask1"
	"newProjectFolders/testTask1/pkg/repository"
)

type TaskInterface interface {
	CreateTask(task tasks.Tasks) (*tasks.Tasks, error)
	GetTaskById(id int) (tasks.Tasks, error)
	UpdateTaskById(id int, name, description, author *string, task tasks.Tasks) error
	DeleteTaskById(id int) error
}

type Service struct {
	TaskInterface
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		TaskInterface: NewTaskService(repository.Task),
	}
}
