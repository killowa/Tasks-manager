package services

import (
	"time"
)

type (
	WorkerPool struct {
		free      map[string]time.Time
		allocated map[string]time.Time
		TTL       time.Duration
		workers   map[string]*Worker
	}
)

var workerPool *WorkerPool = nil

type Pool[T any] interface {
	GetResource() (*T, error)
	ReturnResource(*T) error
	Kill(*T) error
	Expired(*T) bool
	create() (*T, error)
}

func (workerPool *WorkerPool) GetResource() (*Worker, error) {
	if len(workerPool.free) > 0 {
		for workerID := range workerPool.free {
			if workerPool.IsDead(workerID) {
				workerPool.Dispose(workerID)
				delete(workerPool.free, workerID)
				continue
			}

			workerPool.allocated[workerID] = time.Now()
			delete(workerPool.free, workerID)
			return workerPool.workers[workerID], nil
		}
	}

	worker := workerPool.create()

	workerPool.allocated[worker.ID] = time.Now()
	return worker, nil
}

func (workerPool *WorkerPool) Dispose(workerID string) {
	workerPool.workers[workerID].Stop()
	delete(workerPool.workers, workerID)
	delete(workerPool.free, workerID)
	delete(workerPool.allocated, workerID)
}

func (workerPool *WorkerPool) getAllDeathTimesMap() map[string]time.Time {
	deathTimes := make(map[string]time.Time)
	for workerID := range workerPool.free {
		deathTimes[workerID] = workerPool.free[workerID]
	}
	for workerID := range workerPool.allocated {
		deathTimes[workerID] = workerPool.allocated[workerID]
	}

	return deathTimes
}

func (workerPool *WorkerPool) IsDead(workerID string) bool {
	deathTimes := workerPool.getAllDeathTimesMap()
	return time.Now().After(deathTimes[workerID])
}

func GetWorkerPool() *WorkerPool {
	if workerPool == nil {
		workerPool = &WorkerPool{
			free:      make(map[string]time.Time),
			allocated: make(map[string]time.Time),
			TTL:       30 * 60 * time.Second,
			workers:   make(map[string]*Worker),
		}
	}
	return workerPool
}

func (workerPool *WorkerPool) ReturnResource(worker *Worker) error {
	workerPool.free[worker.ID] = time.Now()
	delete(workerPool.allocated, worker.ID)
	return nil
}

func (workerPool *WorkerPool) create() *Worker {
	worker := NewWorker()
	workerPool.workers[worker.ID] = worker
	return worker
}
