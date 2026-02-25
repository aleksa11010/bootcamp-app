package taskmanager

import (
	"testing"
	"time"
)

func TestPriorityString(t *testing.T) {
	tests := []struct {
		name     string
		priority Priority
		want     string
	}{
		{"Low", PriorityLow, "Low"},
		{"Medium", PriorityMedium, "Medium"},
		{"High", PriorityHigh, "High"},
		{"Unknown", Priority(99), "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.priority.String(); got != tt.want {
				t.Errorf("Priority.String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestPriorityValid(t *testing.T) {
	tests := []struct {
		name     string
		priority Priority
		want     bool
	}{
		{"Low is valid", PriorityLow, true},
		{"Medium is valid", PriorityMedium, true},
		{"High is valid", PriorityHigh, true},
		{"Negative is invalid", Priority(-1), false},
		{"OutOfRange is invalid", Priority(99), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.priority.Valid(); got != tt.want {
				t.Errorf("Priority.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatusString(t *testing.T) {
	tests := []struct {
		name   string
		status Status
		want   string
	}{
		{"Pending", StatusPending, "Pending"},
		{"Completed", StatusCompleted, "Completed"},
		{"Unknown", Status(99), "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.String(); got != tt.want {
				t.Errorf("Status.String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestTaskString(t *testing.T) {
	task := Task{
		ID:        1,
		Title:     "Test task",
		Priority:  PriorityHigh,
		Status:    StatusPending,
		CreatedAt: time.Now(),
	}

	got := task.String()
	if got == "" {
		t.Error("Task.String() returned empty string")
	}

	// Verify it contains key info
	for _, substr := range []string{"#1", "Test task", "High", "Pending"} {
		if !contains(got, substr) {
			t.Errorf("Task.String() = %q, expected to contain %q", got, substr)
		}
	}
}

func TestTaskStringCompleted(t *testing.T) {
	now := time.Now()
	task := Task{
		ID:          1,
		Title:       "Done task",
		Priority:    PriorityLow,
		Status:      StatusCompleted,
		CreatedAt:   now.Add(-time.Hour),
		CompletedAt: &now,
	}

	got := task.String()
	if !contains(got, "Completed at") {
		t.Errorf("Task.String() = %q, expected to contain 'Completed at'", got)
	}
}

func TestTaskIsOverdue(t *testing.T) {
	threshold := 24 * time.Hour

	t.Run("pending task past threshold is overdue", func(t *testing.T) {
		task := Task{
			Status:    StatusPending,
			CreatedAt: time.Now().Add(-48 * time.Hour),
		}
		if !task.IsOverdue(threshold) {
			t.Error("expected task to be overdue")
		}
	})

	t.Run("pending task within threshold is not overdue", func(t *testing.T) {
		task := Task{
			Status:    StatusPending,
			CreatedAt: time.Now(),
		}
		if task.IsOverdue(threshold) {
			t.Error("expected task to not be overdue")
		}
	})

	t.Run("completed task is never overdue", func(t *testing.T) {
		task := Task{
			Status:    StatusCompleted,
			CreatedAt: time.Now().Add(-48 * time.Hour),
		}
		if task.IsOverdue(threshold) {
			t.Error("expected completed task to not be overdue")
		}
	})
}

// contains is a simple helper for substring checks.
func contains(s, substr string) bool {
	return len(s) >= len(substr) && searchString(s, substr)
}

func searchString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
