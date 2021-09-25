package garyshker

type Tasks struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Author      string `json:"author" binding:"required"`
}
