package main

import (
	"fmt"
	"log"

	"github.com/sample-go-app/taskmanager"
)

func main() {
	tm := taskmanager.New()

	t1, err := tm.Add("Buy groceries", taskmanager.PriorityMedium)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Added: %s\n", t1)

	t2, err := tm.Add("Write Go tests", taskmanager.PriorityHigh)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Added: %s\n", t2)

	t3, err := tm.Add("Clean the house", taskmanager.PriorityLow)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Added: %s\n", t3)

	fmt.Printf("\nAll tasks (%d):\n", tm.Count())
	for _, t := range tm.List() {
		fmt.Printf("  - %s\n", t)
	}

	if err := tm.Complete(t1.ID); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nCompleted: %s\n", t1.Title)

	fmt.Println("\nPending tasks:")
	for _, t := range tm.ListByStatus(taskmanager.StatusPending) {
		fmt.Printf("  - %s\n", t)
	}

	fmt.Println("\nHigh priority tasks:")
	for _, t := range tm.ListByPriority(taskmanager.PriorityHigh) {
		fmt.Printf("  - %s\n", t)
	}

	stats := tm.Stats()
	fmt.Printf("\nStats: Total=%d, Pending=%d, Completed=%d\n",
		stats.Total, stats.Pending, stats.Completed)
}
