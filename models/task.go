package models

import (
	"context"
	"sync"
	"time"
)

type TaskState int

const (
	TaskPending TaskState = iota
	TaskRunning
	TaskCompleted
	TaskFailed
	TaskCancelled
)

func (ts TaskState) String() string {
	switch ts {
	case TaskPending:
		return "pending"
	case TaskRunning:
		return "running"
	case TaskCompleted:
		return "completed"
	case TaskFailed:
		return "failed"
	case TaskCancelled:
		return "cancelled"
	default:
		return "unknown"
	}
}

// Task represents a schedulable task
type Task struct {
	ID         string
	Name       string
	Interval   time.Duration
	Command    func(context.Context) error
	Context    context.Context
	StartTime  time.Time
	EndTime    *time.Time // Optional end time
	State      TaskState
	LastRun    *time.Time
	NextRun    time.Time
	RunCount   int
	ErrorCount int
	MaxRetries int
	mu         sync.RWMutex
}

// TaskExecution represents the result of a task execution
type TaskExecution struct {
	TaskID    string
	StartTime time.Time
	EndTime   time.Time
	Error     error
	Duration  time.Duration
}
