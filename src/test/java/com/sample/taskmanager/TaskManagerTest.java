package com.sample.taskmanager;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;

import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertThrows;

class TaskManagerTest {

    private TaskManager tm;

    @BeforeEach
    void setUp() {
        tm = new TaskManager();
    }

    @Nested
    @DisplayName("Add operations")
    class AddTests {

        @Test
        @DisplayName("Add creates task with correct fields")
        void addCreatesTask() {
            Task task = tm.add("Test task", Priority.MEDIUM);

            assertEquals(1, task.getId());
            assertEquals("Test task", task.getTitle());
            assertEquals(Priority.MEDIUM, task.getPriority());
            assertEquals(Status.PENDING, task.getStatus());
            assertEquals(1, tm.count());
        }

        @Test
        @DisplayName("Add trims whitespace from title")
        void addTrimsWhitespace() {
            Task task = tm.add("  padded title  ", Priority.LOW);
            assertEquals("padded title", task.getTitle());
        }

        @Test
        @DisplayName("Add auto-increments IDs")
        void addAutoIncrementsIds() {
            Task t1 = tm.add("First", Priority.LOW);
            Task t2 = tm.add("Second", Priority.LOW);
            Task t3 = tm.add("Third", Priority.LOW);

            assertEquals(1, t1.getId());
            assertEquals(2, t2.getId());
            assertEquals(3, t3.getId());
        }

        @Test
        @DisplayName("Add throws on empty title")
        void addThrowsOnEmptyTitle() {
            assertThrows(IllegalArgumentException.class, () -> tm.add("", Priority.LOW));
            assertThrows(IllegalArgumentException.class, () -> tm.add("   ", Priority.LOW));
            assertThrows(IllegalArgumentException.class, () -> tm.add(null, Priority.LOW));
        }

        @Test
        @DisplayName("Add throws on null priority")
        void addThrowsOnNullPriority() {
            assertThrows(IllegalArgumentException.class, () -> tm.add("Valid", null));
        }
    }

    @Nested
    @DisplayName("Get operations")
    class GetTests {

        @Test
        @DisplayName("Get returns correct task")
        void getReturnsTask() {
            Task added = tm.add("Test", Priority.HIGH);
            Task got = tm.get(added.getId());
            assertEquals(added.getId(), got.getId());
        }

        @Test
        @DisplayName("Get throws on invalid ID")
        void getThrowsOnInvalidId() {
            assertThrows(IllegalArgumentException.class, () -> tm.get(0));
            assertThrows(IllegalArgumentException.class, () -> tm.get(-1));
        }

        @Test
        @DisplayName("Get throws on not found")
        void getThrowsOnNotFound() {
            assertThrows(TaskNotFoundException.class, () -> tm.get(999));
        }
    }

    @Nested
    @DisplayName("Complete operations")
    class CompleteTests {

        @Test
        @DisplayName("Complete marks task as completed")
        void completeMarksTask() {
            Task task = tm.add("Complete me", Priority.MEDIUM);
            tm.complete(task.getId());

            Task got = tm.get(task.getId());
            assertEquals(Status.COMPLETED, got.getStatus());
            assertNotNull(got.getCompletedAt());
        }

        @Test
        @DisplayName("Complete throws on already completed")
        void completeThrowsOnAlreadyDone() {
            Task task = tm.add("Done", Priority.LOW);
            tm.complete(task.getId());

            assertThrows(IllegalStateException.class, () -> tm.complete(task.getId()));
        }

        @Test
        @DisplayName("Complete throws on not found")
        void completeThrowsOnNotFound() {
            assertThrows(TaskNotFoundException.class, () -> tm.complete(999));
        }
    }

    @Nested
    @DisplayName("Delete operations")
    class DeleteTests {

        @Test
        @DisplayName("Delete removes task")
        void deleteRemovesTask() {
            Task task = tm.add("Delete me", Priority.LOW);
            tm.delete(task.getId());
            assertEquals(0, tm.count());
        }

        @Test
        @DisplayName("Delete throws on invalid ID")
        void deleteThrowsOnInvalidId() {
            assertThrows(IllegalArgumentException.class, () -> tm.delete(0));
        }

        @Test
        @DisplayName("Delete throws on not found")
        void deleteThrowsOnNotFound() {
            assertThrows(TaskNotFoundException.class, () -> tm.delete(999));
        }
    }

    @Nested
    @DisplayName("List operations")
    class ListTests {

        @Test
        @DisplayName("List returns all tasks in order")
        void listReturnsAllTasks() {
            tm.add("A", Priority.LOW);
            tm.add("B", Priority.MEDIUM);
            tm.add("C", Priority.HIGH);

            List<Task> tasks = tm.list();
            assertEquals(3, tasks.size());
            assertEquals(1, tasks.get(0).getId());
            assertEquals(2, tasks.get(1).getId());
            assertEquals(3, tasks.get(2).getId());
        }

        @Test
        @DisplayName("ListByStatus filters correctly")
        void listByStatusFilters() {
            Task t1 = tm.add("Pending", Priority.LOW);
            tm.add("Also pending", Priority.MEDIUM);
            tm.complete(t1.getId());

            assertEquals(1, tm.listByStatus(Status.PENDING).size());
            assertEquals(1, tm.listByStatus(Status.COMPLETED).size());
        }

        @Test
        @DisplayName("ListByPriority filters correctly")
        void listByPriorityFilters() {
            tm.add("Low task", Priority.LOW);
            tm.add("High task 1", Priority.HIGH);
            tm.add("High task 2", Priority.HIGH);

            assertEquals(2, tm.listByPriority(Priority.HIGH).size());
            assertEquals(1, tm.listByPriority(Priority.LOW).size());
            assertEquals(0, tm.listByPriority(Priority.MEDIUM).size());
        }
    }

    @Nested
    @DisplayName("Stats operations")
    class StatsTests {

        @Test
        @DisplayName("Stats returns correct counts")
        void statsReturnsCorrectCounts() {
            Task t1 = tm.add("Task 1", Priority.LOW);
            tm.add("Task 2", Priority.MEDIUM);
            tm.add("Task 3", Priority.HIGH);
            tm.complete(t1.getId());

            Stats stats = tm.stats();
            assertEquals(3, stats.getTotal());
            assertEquals(2, stats.getPending());
            assertEquals(1, stats.getCompleted());
        }

        @Test
        @DisplayName("Stats returns zero for empty manager")
        void statsReturnsZeroForEmpty() {
            Stats stats = tm.stats();

            assertEquals(0, stats.getTotal());
            assertEquals(0, stats.getPending());
            assertEquals(0, stats.getCompleted());
        }
    }
}
