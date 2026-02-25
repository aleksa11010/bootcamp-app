package com.sample.taskmanager;

public class App {

    public static void main(String[] args) {
        TaskManager tm = new TaskManager();

        Task t1 = tm.add("Buy groceries", Priority.MEDIUM);
        System.out.println("Added: " + t1);

        Task t2 = tm.add("Write Java tests", Priority.HIGH);
        System.out.println("Added: " + t2);

        Task t3 = tm.add("Clean the house", Priority.LOW);
        System.out.println("Added: " + t3);

        System.out.printf("%nAll tasks (%d):%n", tm.count());
        for (Task t : tm.list()) {
            System.out.println("  - " + t);
        }

        tm.complete(t1.getId());
        System.out.printf("%nCompleted: %s%n", t1.getTitle());

        System.out.println("\nPending tasks:");
        for (Task t : tm.listByStatus(Status.PENDING)) {
            System.out.println("  - " + t);
        }

        System.out.println("\nHigh priority tasks:");
        for (Task t : tm.listByPriority(Priority.HIGH)) {
            System.out.println("  - " + t);
        }

        Stats stats = tm.stats();
        System.out.printf("%nStats: Total=%d, Pending=%d, Completed=%d%n",
                stats.getTotal(), stats.getPending(), stats.getCompleted());
    }
}
