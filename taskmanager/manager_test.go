package taskmanager

import (
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	tm := New()
	if tm == nil {
		t.Fatal("New() returned nil")
	}
	if tm.Count() != 0 {
		t.Errorf("New manager should have 0 tasks, got %d", tm.Count())
	}
}

func TestAdd(t *testing.T) {
	tm := New()

	task, err := tm.Add("Test task", PriorityMedium)
	if err != nil {
		t.Fatalf("Add() returned error: %v", err)
	}
	if task.ID != 1 {
		t.Errorf("first task ID = %d, want 1", task.ID)
	}
	if task.Title != "Test task" {
		t.Errorf("task.Title = %q, want %q", task.Title, "Test task")
	}
	if task.Priority != PriorityMedium {
		t.Errorf("task.Priority = %v, want %v", task.Priority, PriorityMedium)
	}
	if task.Status != StatusPending {
		t.Errorf("task.Status = %v, want %v", task.Status, StatusPending)
	}
	if tm.Count() != 1 {
		t.Errorf("Count() = %d, want 1", tm.Count())
	}
}

func TestAddEmptyTitle(t *testing.T) {
	tm := New()

	_, err := tm.Add("", PriorityLow)
	if err != ErrEmptyTitle {
		t.Errorf("Add('') error = %v, want %v", err, ErrEmptyTitle)
	}

	_, err = tm.Add("   ", PriorityLow)
	if err != ErrEmptyTitle {
		t.Errorf("Add('   ') error = %v, want %v", err, ErrEmptyTitle)
	}
}

func TestAddInvalidPriority(t *testing.T) {
	tm := New()

	_, err := tm.Add("Valid title", Priority(99))
	if err != ErrInvalidPriority {
		t.Errorf("Add() with invalid priority error = %v, want %v", err, ErrInvalidPriority)
	}
}

func TestAddTrimsWhitespace(t *testing.T) {
	tm := New()

	task, err := tm.Add("  padded title  ", PriorityLow)
	if err != nil {
		t.Fatalf("Add() returned error: %v", err)
	}
	if task.Title != "padded title" {
		t.Errorf("task.Title = %q, want %q", task.Title, "padded title")
	}
}

func TestAddAutoIncrementIDs(t *testing.T) {
	tm := New()

	t1, _ := tm.Add("First", PriorityLow)
	t2, _ := tm.Add("Second", PriorityLow)
	t3, _ := tm.Add("Third", PriorityLow)

	if t1.ID != 1 || t2.ID != 2 || t3.ID != 3 {
		t.Errorf("IDs = (%d, %d, %d), want (1, 2, 3)", t1.ID, t2.ID, t3.ID)
	}
}

func TestGet(t *testing.T) {
	tm := New()
	added, _ := tm.Add("Test", PriorityHigh)

	got, err := tm.Get(added.ID)
	if err != nil {
		t.Fatalf("Get() returned error: %v", err)
	}
	if got.ID != added.ID {
		t.Errorf("Get().ID = %d, want %d", got.ID, added.ID)
	}
}

func TestGetInvalidID(t *testing.T) {
	tm := New()

	_, err := tm.Get(0)
	if err != ErrInvalidID {
		t.Errorf("Get(0) error = %v, want %v", err, ErrInvalidID)
	}

	_, err = tm.Get(-1)
	if err != ErrInvalidID {
		t.Errorf("Get(-1) error = %v, want %v", err, ErrInvalidID)
	}
}

func TestGetNotFound(t *testing.T) {
	tm := New()

	_, err := tm.Get(999)
	if err != ErrNotFound {
		t.Errorf("Get(999) error = %v, want %v", err, ErrNotFound)
	}
}

func TestComplete(t *testing.T) {
	tm := New()
	task, _ := tm.Add("Complete me", PriorityMedium)

	err := tm.Complete(task.ID)
	if err != nil {
		t.Fatalf("Complete() returned error: %v", err)
	}

	got, _ := tm.Get(task.ID)
	if got.Status != StatusCompleted {
		t.Errorf("task.Status = %v, want %v", got.Status, StatusCompleted)
	}
	if got.CompletedAt == nil {
		t.Error("task.CompletedAt should not be nil after completion")
	}
}

func TestCompleteAlreadyDone(t *testing.T) {
	tm := New()
	task, _ := tm.Add("Done", PriorityLow)
	_ = tm.Complete(task.ID)

	err := tm.Complete(task.ID)
	if err != ErrAlreadyDone {
		t.Errorf("Complete() twice error = %v, want %v", err, ErrAlreadyDone)
	}
}

func TestCompleteInvalidID(t *testing.T) {
	tm := New()

	err := tm.Complete(0)
	if err != ErrInvalidID {
		t.Errorf("Complete(0) error = %v, want %v", err, ErrInvalidID)
	}
}

func TestCompleteNotFound(t *testing.T) {
	tm := New()

	err := tm.Complete(999)
	if err != ErrNotFound {
		t.Errorf("Complete(999) error = %v, want %v", err, ErrNotFound)
	}
}

func TestDelete(t *testing.T) {
	tm := New()
	task, _ := tm.Add("Delete me", PriorityLow)

	err := tm.Delete(task.ID)
	if err != nil {
		t.Fatalf("Delete() returned error: %v", err)
	}
	if tm.Count() != 0 {
		t.Errorf("Count() after delete = %d, want 0", tm.Count())
	}
}

func TestDeleteInvalidID(t *testing.T) {
	tm := New()

	err := tm.Delete(0)
	if err != ErrInvalidID {
		t.Errorf("Delete(0) error = %v, want %v", err, ErrInvalidID)
	}
}

func TestDeleteNotFound(t *testing.T) {
	tm := New()

	err := tm.Delete(999)
	if err != ErrNotFound {
		t.Errorf("Delete(999) error = %v, want %v", err, ErrNotFound)
	}
}

func TestList(t *testing.T) {
	tm := New()
	tm.Add("A", PriorityLow)
	tm.Add("B", PriorityMedium)
	tm.Add("C", PriorityHigh)

	tasks := tm.List()
	if len(tasks) != 3 {
		t.Fatalf("List() returned %d tasks, want 3", len(tasks))
	}

	// Verify ordering by ID
	for i, task := range tasks {
		if task.ID != i+1 {
			t.Errorf("tasks[%d].ID = %d, want %d", i, task.ID, i+1)
		}
	}
}

func TestListByStatus(t *testing.T) {
	tm := New()
	t1, _ := tm.Add("Pending", PriorityLow)
	tm.Add("Also pending", PriorityMedium)
	tm.Complete(t1.ID)

	pending := tm.ListByStatus(StatusPending)
	if len(pending) != 1 {
		t.Errorf("ListByStatus(Pending) returned %d tasks, want 1", len(pending))
	}

	completed := tm.ListByStatus(StatusCompleted)
	if len(completed) != 1 {
		t.Errorf("ListByStatus(Completed) returned %d tasks, want 1", len(completed))
	}
}

func TestListByPriority(t *testing.T) {
	tm := New()
	tm.Add("Low task", PriorityLow)
	tm.Add("High task 1", PriorityHigh)
	tm.Add("High task 2", PriorityHigh)

	high := tm.ListByPriority(PriorityHigh)
	if len(high) != 2 {
		t.Errorf("ListByPriority(High) returned %d tasks, want 2", len(high))
	}

	low := tm.ListByPriority(PriorityLow)
	if len(low) != 1 {
		t.Errorf("ListByPriority(Low) returned %d tasks, want 1", len(low))
	}

	medium := tm.ListByPriority(PriorityMedium)
	if len(medium) != 0 {
		t.Errorf("ListByPriority(Medium) returned %d tasks, want 0", len(medium))
	}
}

func TestStats(t *testing.T) {
	tm := New()
	t1, _ := tm.Add("Task 1", PriorityLow)
	tm.Add("Task 2", PriorityMedium)
	tm.Add("Task 3", PriorityHigh)
	tm.Complete(t1.ID)

	stats := tm.Stats()
	if stats.Total != 3 {
		t.Errorf("stats.Total = %d, want 3", stats.Total)
	}
	if stats.Pending != 2 {
		t.Errorf("stats.Pending = %d, want 2", stats.Pending)
	}
	if stats.Completed != 1 {
		t.Errorf("stats.Completed = %d, want 1", stats.Completed)
	}
}

func TestConcurrentAccess(t *testing.T) {
	tm := New()
	var wg sync.WaitGroup
	numGoroutines := 100

	// Concurrent adds
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(n int) {
			defer wg.Done()
			_, err := tm.Add("Task", PriorityMedium)
			if err != nil {
				t.Errorf("concurrent Add() failed: %v", err)
			}
		}(i)
	}
	wg.Wait()

	if tm.Count() != numGoroutines {
		t.Errorf("Count() = %d after %d concurrent adds", tm.Count(), numGoroutines)
	}

	// Concurrent reads
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			_ = tm.List()
			_ = tm.Count()
			_ = tm.Stats()
		}()
	}
	wg.Wait()
}
