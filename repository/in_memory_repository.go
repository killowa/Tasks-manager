package repository

import (
	"fmt"
	"sync"
	"tasks_manager/models"

	"github.com/google/uuid"
)

var InMemoryRepo *InMemoryRepository

// InMemoryRepository implements TaskRepository
type InMemoryRepository struct {
	tasks map[string]*models.Task
	mu    sync.RWMutex
}

// NewInMemoryRepository creates a new in-memory repository
func GetInMemoryRepository() *InMemoryRepository {
	if InMemoryRepo == nil {
		InMemoryRepo = &InMemoryRepository{
			tasks: make(map[string]*models.Task),
		}
	}
	return InMemoryRepo
}

// Store saves a task to the repository
func (r *InMemoryRepository) Store(task *models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	task.ID = uuid.New().String()
	r.tasks[task.ID] = task
	return nil
}

// Get retrieves a task by ID
func (r *InMemoryRepository) Get(id string) (*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	task, exists := r.tasks[id]
	if !exists {
		return nil, fmt.Errorf("task with id %s not found", id)
	}
	return task, nil
}

// List returns all tasks
func (r *InMemoryRepository) List() ([]*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]*models.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// Delete removes a task from the repository
func (r *InMemoryRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return fmt.Errorf("task with id %s not found", id)
	}
	delete(r.tasks, id)
	return nil
}
