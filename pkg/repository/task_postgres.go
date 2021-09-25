package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	tasks "newProjectFolders/testTask1"
)

type TaskStruct struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskStruct {
	return &TaskStruct{db: db}
}

func (t *TaskStruct) CreateTask(task tasks.Tasks) (*tasks.Tasks, error) {
	printedTask := &tasks.Tasks{}

	query := fmt.Sprintf("insert into %s (name, description, author) values ($1,$2,$3) returning *", "task_info")
	row := t.db.QueryRow(query, task.Name, task.Description, task.Author)
	if err := row.Scan(&printedTask.Id, &printedTask.Name, &printedTask.Description, &printedTask.Author); err != nil {
		return nil, err
	}
	return printedTask, nil
}

func (t *TaskStruct) GetTaskById(id int) (tasks.Tasks, error) {
	var task tasks.Tasks
	query := fmt.Sprintf("select * from %s where id=$1", "task_info")
	err := t.db.Get(&task, query, id)
	return task, err
}

func (t *TaskStruct) UpdateTaskById(id int, task tasks.Tasks) error {
	query := fmt.Sprintf("update %s set name=$1,  description=$2, author=$3 where id = $4", "task_info")
	_, err := t.db.Query(query, task.Name, task.Description, task.Author, id)
	return err
}

func (t *TaskStruct) DeleteTaskById(id int) error {
	query := fmt.Sprintf("delete from %s where id=$1", "task_info")
	_, err := t.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
