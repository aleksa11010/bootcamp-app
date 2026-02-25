package com.sample.taskmanager;

import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.time.Duration;

public class Task {

    private static final DateTimeFormatter FORMATTER = DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm");

    private final int id;
    private final String title;
    private final Priority priority;
    private Status status;
    private final LocalDateTime createdAt;
    private LocalDateTime completedAt;

    public Task(int id, String title, Priority priority) {
        this.id = id;
        this.title = title;
        this.priority = priority;
        this.status = Status.PENDING;
        this.createdAt = LocalDateTime.now();
    }

    public int getId() {
        return id;
    }

    public String getTitle() {
        return title;
    }

    public Priority getPriority() {
        return priority;
    }

    public Status getStatus() {
        return status;
    }

    public LocalDateTime getCreatedAt() {
        return createdAt;
    }

    public LocalDateTime getCompletedAt() {
        return completedAt;
    }

    void markCompleted() {
        this.status = Status.COMPLETED;
        this.completedAt = LocalDateTime.now();
    }

    public boolean isOverdue(Duration threshold) {
        if (status == Status.COMPLETED) {
            return false;
        }
        return Duration.between(createdAt, LocalDateTime.now()).compareTo(threshold) > 0;
    }

    @Override
    public String toString() {
        String statusStr = (completedAt != null)
                ? "Completed at " + completedAt.format(FORMATTER)
                : status.toString();
        return String.format("[#%d] %s (Priority: %s, Status: %s)", id, title, priority, statusStr);
    }
}
