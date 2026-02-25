package taskmanager

import (
	"errors"
	"strings"
	"sync"
	"time"
)

var (
	ErrEmptyTitle    = errors.New("task title cannot be empty")
	ErrInvalidID     = errors.New("invalid task ID")
	ErrNotFound      = errors.New("task not found")
	ErrAlreadyDone   = errors.New("task is already completed")
	ErrInvalidPriority = errors.New("invalid priority value")
)

// Stats holds aggregate statistics about tasks.
type Stats struct {
	Total     int
	Pending   int
	Completed int
}

// Manager manages a collection of tasks with thread-safe operations.
type Manager struct {
	mu     sync.RWMutex
	tasks  map[int]*Task
	nextID int
}

// New creates a new Manager instance.
func New() *Manager {
	return &Manager{
		tasks:  make(map[int]*Task),
		nextID: 1,
	}
}

// Add creates a new task with the given title and priority.
func (m *Manager) Add(title string, priority Priority) (*Task, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return nil, ErrEmptyTitle
	}
	if !priority.Valid() {
		return nil, ErrInvalidPriority
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	task := &Task{
		ID:        m.nextID,
		Title:     title,
		Priority:  priority,
		Status:    StatusPending,
		CreatedAt: time.Now(),
	}
	m.tasks[m.nextID] = task
	m.nextID++

	return task, nil
}

// Get retrieves a task by its ID.
func (m *Manager) Get(id int) (*Task, error) {
	if id <= 0 {
		return nil, ErrInvalidID
	}

	m.mu.RLock()
	defer m.mu.RUnlock()

	task, ok := m.tasks[id]
	if !ok {
		return nil, ErrNotFound
	}
	return task, nil
}

// Complete marks a task as completed.
func (m *Manager) Complete(id int) error {
	if id <= 0 {
		return ErrInvalidID
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	task, ok := m.tasks[id]
	if !ok {
		return ErrNotFound
	}
	if task.Status == StatusCompleted {
		return ErrAlreadyDone
	}

	now := time.Now()
	task.Status = StatusCompleted
	task.CompletedAt = &now
	return nil
}

// Delete removes a task by its ID.
func (m *Manager) Delete(id int) error {
	if id <= 0 {
		return ErrInvalidID
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.tasks[id]; !ok {
		return ErrNotFound
	}
	delete(m.tasks, id)
	return nil
}

// List returns all tasks sorted by ID.
func (m *Manager) List() []*Task {
	m.mu.RLock()
	defer m.mu.RUnlock()

	tasks := make([]*Task, 0, len(m.tasks))
	for i := 1; i < m.nextID; i++ {
		if t, ok := m.tasks[i]; ok {
			tasks = append(tasks, t)
		}
	}
	return tasks
}

// ListByStatus returns tasks filtered by status.
func (m *Manager) ListByStatus(status Status) []*Task {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var tasks []*Task
	for i := 1; i < m.nextID; i++ {
		if t, ok := m.tasks[i]; ok && t.Status == status {
			tasks = append(tasks, t)
		}
	}
	return tasks
}

// ListByPriority returns tasks filtered by priority.
func (m *Manager) ListByPriority(priority Priority) []*Task {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var tasks []*Task
	for i := 1; i < m.nextID; i++ {
		if t, ok := m.tasks[i]; ok && t.Priority == priority {
			tasks = append(tasks, t)
		}
	}
	return tasks
}

// Count returns the total number of tasks.
func (m *Manager) Count() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.tasks)
}

// Stats returns aggregate statistics.
func (m *Manager) Stats() Stats {
	m.mu.RLock()
	defer m.mu.RUnlock()

	s := Stats{Total: len(m.tasks)}
	for _, t := range m.tasks {
		switch t.Status {
		case StatusPending:
			s.Pending++
		case StatusCompleted:
			s.Completed++
		}
	}
	return s
}
