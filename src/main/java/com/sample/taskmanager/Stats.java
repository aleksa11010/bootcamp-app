package com.sample.taskmanager;

public class Stats {

    private final int total;
    private final int pending;
    private final int completed;

    public Stats(int total, int pending, int completed) {
        this.total = total;
        this.pending = pending;
        this.completed = completed;
    }

    public int getTotal() {
        return total;
    }

    public int getPending() {
        return pending;
    }

    public int getCompleted() {
        return completed;
    }

    @Override
    public String toString() {
        return String.format("Stats{total=%d, pending=%d, completed=%d}", total, pending, completed);
    }
}
