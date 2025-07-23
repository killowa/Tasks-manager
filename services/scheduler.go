package services

import (
	"fmt"
	"tasks_manager/models"
	"tasks_manager/repository"
	"time"
)

// Scheduler
// fire a watcher routine
// Maintain a list of scheduled tasks with in the current hr
// refresh this list every 1 hr
// check for available workers
// assign the next scheduled task for the available worker

// Worker
// Keep checking queue for any tasks
// poll from queue and execute
// worker expect command fn with it's args

type Scheduler struct{}

var CurrentExecutionSpan = time.Now().Add(60 * 1000 * 1000)

func GetScheduler() *Scheduler {
	return &Scheduler{}
}

func (sch *Scheduler) Schedule(task *models.Task, startTime time.Time) {
	task.StartTime = startTime
	workerPool := GetWorkerPool()
	if startTime.Before(CurrentExecutionSpan) {
		worker, err := workerPool.GetResource()
		if err != nil {
			return
		}
		worker.Take(task)
		worker.Start()
		workerPool.ReturnResource(worker)
		 fmt.Println("Worker returned", worker.ID)
	}
	repository.GetInMemoryRepository().Store(task)
}
