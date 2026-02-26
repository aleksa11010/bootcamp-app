package com.sample.taskmanager;

import java.util.ArrayList;
import java.util.Collections;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

public class TaskManager {

    private static final int MIN_VALID_ID = 1;

    private final Map<Integer, Task> tasks = new LinkedHashMap<>();
    private int nextId = 1;

    public Task add(String title, Priority priority) {
        if (title == null || title.trim().isEmpty()) {
            throw new IllegalArgumentException("Task title cannot be empty");
        }
        if (priority == null) {
            throw new IllegalArgumentException("Priority cannot be null");
        }

        String trimmed = title.trim();
        Task task = new Task(nextId, trimmed, priority);
        tasks.put(nextId, task);
        nextId++;
        return task;
    }

    public Task get(int id) {
        if (id < MIN_VALID_ID) {
            throw new IllegalArgumentException("Invalid task ID");
        }
        Task task = tasks.get(id);
        if (task == null) {
            throw new TaskNotFoundException("Task not found: " + id);
        }
        return task;
    }

    public void complete(int id) {
        Task task = get(id);
        if (task.getStatus() == Status.COMPLETED) {
            throw new IllegalStateException("Task is already completed");
        }
        task.markCompleted();
    }

    public void delete(int id) {
        if (id < MIN_VALID_ID) {
            throw new IllegalArgumentException("Invalid task ID");
        }
        if (tasks.remove(id) == null) {
            throw new TaskNotFoundException("Task not found: " + id);
        }
    }

    public List<Task> list() {
        return Collections.unmodifiableList(new ArrayList<>(tasks.values()));
    }

    public List<Task> listByStatus(Status status) {
        return tasks.values().stream()
                .filter(t -> t.getStatus() == status)
                .collect(Collectors.toList());
    }

    public List<Task> listByPriority(Priority priority) {
        return tasks.values().stream()
                .filter(t -> t.getPriority() == priority)
                .collect(Collectors.toList());
    }

    public int count() {
        return tasks.size();
    }

    public Stats stats() {
        int pending = 0;
        int completed = 0;
        for (Task t : tasks.values()) {
            if (t.getStatus() == Status.PENDING) {
                pending++;
            } else if (t.getStatus() == Status.COMPLETED) {
                completed++;
            }
        }
        return new Stats(tasks.size(), pending, completed);
    }
}
