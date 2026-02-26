package com.sample.taskmanager;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.EnumSource;

import java.time.Duration;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertNull;
import static org.junit.jupiter.api.Assertions.assertTrue;

class TaskTest {

    @Test
    @DisplayName("New task has correct defaults")
    void newTaskHasCorrectDefaults() {
        Task task = new Task(1, "Test task", Priority.HIGH);

        assertEquals(1, task.getId());
        assertEquals("Test task", task.getTitle());
        assertEquals(Priority.HIGH, task.getPriority());
        assertEquals(Status.PENDING, task.getStatus());
        assertNotNull(task.getCreatedAt());
        assertNull(task.getCompletedAt());
    }

    @ParameterizedTest
    @EnumSource(Priority.class)
    @DisplayName("Priority enum has correct labels")
    void priorityLabels(Priority priority) {
        assertNotNull(priority.getLabel());
        assertFalse(priority.getLabel().isEmpty());
        assertEquals(priority.getLabel(), priority.toString());
    }

    @ParameterizedTest
    @EnumSource(Status.class)
    @DisplayName("Status enum has correct labels")
    void statusLabels(Status status) {
        assertNotNull(status.getLabel());
        assertFalse(status.getLabel().isEmpty());
        assertEquals(status.getLabel(), status.toString());
    }

    @Test
    @DisplayName("toString contains key information")
    void toStringContainsKeyInfo() {
        Task task = new Task(42, "My task", Priority.MEDIUM);
        String result = task.toString();

        assertTrue(result.contains("#42"));
        assertTrue(result.contains("My task"));
        assertTrue(result.contains("Medium"));
        assertTrue(result.contains("Pending"));
    }

    @Test
    @DisplayName("toString shows completion time when completed")
    void toStringShowsCompletionTime() {
        Task task = new Task(1, "Done task", Priority.LOW);
        task.markCompleted();

        String result = task.toString();
        assertTrue(result.contains("Completed at"));
    }

    @Test
    @DisplayName("Pending task past threshold is overdue")
    void pendingTaskPastThresholdIsOverdue() {
        Task task = new Task(1, "Old task", Priority.HIGH);
        // A task just created cannot be overdue with a 0-second threshold... unless we use negative
        assertTrue(task.isOverdue(Duration.ZERO));
    }

    @Test
    @DisplayName("Completed task is never overdue")
    void completedTaskIsNeverOverdue() {
        Task task = new Task(1, "Done", Priority.LOW);
        task.markCompleted();

        assertFalse(task.isOverdue(Duration.ZERO));
    }
}
