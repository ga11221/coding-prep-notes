package main

import "container/heap"

type Edge struct {
	to     int
	weight int
}

type Item struct {
	node     int
	distance int
	index    int
}

// implement container/heap.Interface for PriorityQueue
// Less: lower distance has higher priority
// Push/Pop: standard append/slice shrink

func dijkstra(graph [][]Edge, start int) []int {
	// 1. init dist array with max value, set dist[start]=0
	// 2. init min-heap, push start with distance 0
	// 3. while heap not empty:
	//    a. pop min item
	//    b. if stale (dist > known best), skip
	//    c. for each neighbor: if shorter path found, update dist + push
	// 4. return dist
	return nil
}
