package main

// Partition tasks by type. Append to queue, update window length.
// If window length < n+1, append task from next partition.
// If no next partition → idle.
// Always pick the most frequent available task to minimize idle time.
//
// Formula: maxFreq, idleSlots = (maxFreq-1)*(n+1).
// result = max(len(tasks), idleSlots + count of tasks with frequency == maxFreq).

func leastInterval(tasks []byte, n int) int {
	return 0
}

func leastIntervalFormula(tasks []byte, n int) int {
	return 0
}
