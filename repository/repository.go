package repository

import "tasks_manager/models"

type TaskRepository interface {
	Store(task *models.Task) error
	Get(id string) (*models.Task, error)
	List() ([]*models.Task, error)
	Delete(id string) error
}
