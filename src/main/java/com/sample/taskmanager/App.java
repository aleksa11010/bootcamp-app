package com.sample.taskmanager;

import java.util.logging.Logger;

public final class App {

    private static final Logger LOGGER = Logger.getLogger(App.class.getName());

    private App() {
        // Utility class
    }

    public static void main(String[] args) {
        TaskManager tm = new TaskManager();

        Task t1 = tm.add("Buy groceries", Priority.MEDIUM);
        LOGGER.info("Added: " + t1);

        Task t2 = tm.add("Write Java tests", Priority.HIGH);
        LOGGER.info("Added: " + t2);

        Task t3 = tm.add("Clean the house", Priority.LOW);
        LOGGER.info("Added: " + t3);

        LOGGER.info(String.format("%nAll tasks (%d):", tm.count()));
        for (Task t : tm.list()) {
            LOGGER.info("  - " + t);
        }

        tm.complete(t1.getId());
        LOGGER.info(String.format("%nCompleted: %s", t1.getTitle()));

        LOGGER.info("\nPending tasks:");
        for (Task t : tm.listByStatus(Status.PENDING)) {
            LOGGER.info("  - " + t);
        }

        LOGGER.info("\nHigh priority tasks:");
        for (Task t : tm.listByPriority(Priority.HIGH)) {
            LOGGER.info("  - " + t);
        }

        Stats stats = tm.stats();
        LOGGER.info(String.format("%nStats: Total=%d, Pending=%d, Completed=%d",
                stats.getTotal(), stats.getPending(), stats.getCompleted()));
    }
}
