package services

import (
	"context"
	"fmt"
	"tasks_manager/models"

	"github.com/google/uuid"
)

type Worker struct {
	ID         string
	ctx        context.Context
	cancel     context.CancelFunc
	tasksQueue chan *models.Task
}

func (w *Worker) Stop() {
	w.cancel()
}

func NewWorker() *Worker {
	ctx, cancel := context.WithCancel(context.Background())
	return &Worker{
		ID:         uuid.New().String(),
		ctx:        ctx,
		cancel:     cancel,
		tasksQueue: make(chan *models.Task, 10),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			select {
			case <-w.ctx.Done():
				fmt.Println("worker stopped")
				return
			case task := <-w.tasksQueue:
				err := task.Command(task.Context)
				if err != nil {
					fmt.Printf("Task %s failed: %v\n", task.Name, err)
				} else {
					fmt.Printf("Task %s completed successfully\n", task.Name)
				}
			}
		}
	}()
}
func (w *Worker) Take(task *models.Task) error {
	w.tasksQueue <- task
	return nil
}
