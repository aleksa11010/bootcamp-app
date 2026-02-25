package taskmanager

import (
	"fmt"
	"time"
)

// Priority represents the urgency level of a task.
type Priority int

const (
	PriorityLow    Priority = iota
	PriorityMedium
	PriorityHigh
)

// String returns the human-readable name for a Priority.
func (p Priority) String() string {
	switch p {
	case PriorityLow:
		return "Low"
	case PriorityMedium:
		return "Medium"
	case PriorityHigh:
		return "High"
	default:
		return "Unknown"
	}
}

// Valid returns true if the priority is a recognized value.
func (p Priority) Valid() bool {
	return p >= PriorityLow && p <= PriorityHigh
}

// Status represents the current state of a task.
type Status int

const (
	StatusPending   Status = iota
	StatusCompleted
)

// String returns the human-readable name for a Status.
func (s Status) String() string {
	switch s {
	case StatusPending:
		return "Pending"
	case StatusCompleted:
		return "Completed"
	default:
		return "Unknown"
	}
}

// Task represents a single to-do item.
type Task struct {
	ID          int
	Title       string
	Priority    Priority
	Status      Status
	CreatedAt   time.Time
	CompletedAt *time.Time
}

// String returns a formatted representation of the task.
func (t Task) String() string {
	status := t.Status.String()
	if t.CompletedAt != nil {
		status = fmt.Sprintf("Completed at %s", t.CompletedAt.Format(time.RFC822))
	}
	return fmt.Sprintf("[#%d] %s (Priority: %s, Status: %s)", t.ID, t.Title, t.Priority, status)
}

// IsOverdue checks if a pending task was created more than the given duration ago.
func (t Task) IsOverdue(threshold time.Duration) bool {
	if t.Status == StatusCompleted {
		return false
	}
	return time.Since(t.CreatedAt) > threshold
}
