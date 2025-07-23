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

type Task struct {
	ID        string
	Name      string
	Command   func(context.Context) error
	Context   context.Context
	StartTime time.Time
	State     TaskState
	mu        sync.RWMutex
}


