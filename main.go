package main

import (
	"context"
	"fmt"
	"tasks_manager/models"
	"tasks_manager/services"
	"time"
)

func sampleTask(ctx context.Context) error {
	fmt.Println("Executing sample task...")
	time.Sleep(2 * time.Second)
	return nil
}

func main() {
	fmt.Println("=== Task Manager Demo ===")

	// Create repository
	// repo := repository.GetInMemoryRepository()

	// Create a sample task
	task := &models.Task{
		Name:       "Sample Task",
		Interval:   30 * time.Second,
		Command:    sampleTask,
		Context:    context.Background(),
		StartTime:  time.Now(),
		State:      models.TaskPending,
		NextRun:    time.Now().Add(5 * time.Second),
		MaxRetries: 3,
	}
	services.GetScheduler().Schedule(task, time.Now())

	// Wait for the task to be processed
	fmt.Println("Waiting for task to be processed...")
	time.Sleep(5 * time.Second)

	// Store the task
	// 	if err := repo.Store(task); err != nil {
	// 		log.Fatalf("Failed to store task: %v", err)
	// 	}
	//
	// 	fmt.Printf("Created task: %s (ID: %s)\n", task.Name, task.ID)
	//
	// 	// List all tasks
	// 	tasks, err := repo.List()
	// 	if err != nil {
	// 		log.Fatalf("Failed to list tasks: %v", err)
	// 	}
	//
	// 	fmt.Printf("\nFound %d tasks:\n", len(tasks))
	// 	for _, t := range tasks {
	// 		fmt.Printf("- %s: %s (Next run: %s)\n",
	// 			t.Name, t.State.String(), t.NextRun.Format("15:04:05"))
	// 	}
	//
	// 	// Get specific task
	// 	retrieved, err := repo.Get(task.ID)
	// 	if err != nil {
	// 		log.Fatalf("Failed to get task: %v", err)
	// 	}
	//
	// 	fmt.Printf("\nRetrieved task: %s\n", retrieved.Name)
	//
	// 	// Update task state
	// 	retrieved.State = models.TaskRunning
	// 	retrieved.RunCount++
	//
	// 	if err := repo.Store(retrieved); err != nil {
	// 		log.Fatalf("Failed to update task: %v", err)
	// 	}
	//
	// 	fmt.Printf("Updated task state to: %s\n", retrieved.State.String())
	//
	// 	// Delete task
	// 	if err := repo.Delete(task.ID); err != nil {
	// 		log.Fatalf("Failed to delete task: %v", err)
	// 	}
	//
	// 	fmt.Printf("Deleted task: %s\n", task.Name)
	//
	// 	// Verify deletion
	// 	remainingTasks, err := repo.List()
	// 	if err != nil {
	// 		log.Fatalf("Failed to list remaining tasks: %v", err)
	// 	}
	//
	// 	fmt.Printf("\nRemaining tasks: %d\n", len(remainingTasks))
	// 	fmt.Println("Demo completed successfully!")
}
