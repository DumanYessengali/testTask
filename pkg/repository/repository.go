package repository

import (
	"github.com/jmoiron/sqlx"
	tasks "newProjectFolders/testTask1"
)

type Task interface {
	CreateTask(task tasks.Tasks) (*tasks.Tasks, error)
	GetTaskById(id int) (tasks.Tasks, error)
	UpdateTaskById(id int, task tasks.Tasks) error
	DeleteTaskById(id int) error
}

type Repository struct {
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Task: NewTaskPostgres(db),
	}
}
