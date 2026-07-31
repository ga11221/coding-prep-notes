package main

import (
	"container/heap"
	"math"
)

type Edge struct {
	to     int
	weight int
}

type Item struct {
	node     int
	distance int
}

// implement container/heap.Interface for PriorityQueue
// Less: lower distance has higher priority
// Push/Pop: standard append/slice shrink

type PriorityQueue struct {
	items []*Item
}

func (pq PriorityQueue) Len() int {
	return len(pq.items)
}

func (pq *PriorityQueue) Push(x any) {
	pq.items = append(pq.items, x.(*Item))
}

func (pq *PriorityQueue) Pop() any {
	items := pq.items
	length := len(items)
	if length > 0 {
		item := items[length-1]
		pq.items = pq.items[:length-1]
		return item
	}
	return nil
}

func (pq PriorityQueue) Less(i, j int) bool {
	items := pq.items
	return items[i].distance < items[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	items := pq.items
	tmp := items[i]
	items[i] = items[j]
	items[j] = tmp
}

func dijkstra(graph [][]Edge, start int) []int {
	// 1. init dist array with max value, set items[start]=0
	dist := make([]int, len(graph))
	for i := range graph {
		dist[i] = math.MaxInt
	}
	dist[start] = 0

	pq := &PriorityQueue{
		items: []*Item{{node: start, distance: 0}},
	}
	// 2. init min-heap, push start with distance 0
	heap.Init(pq)

	// 3. while heap not empty:
	for pq.Len() > 0 {
		//    a. pop min item
		item := heap.Pop(pq).(*Item)
		//    b. if stale (items > known best), skip
		if item.distance > dist[item.node] {
			continue
		}
		//    c. for each neighbor: if shorter path found, update items + push
		for _, e := range graph[item.node] {
			if nd := item.distance + e.weight; nd < dist[e.to] {
				dist[e.to] = nd
				heap.Push(pq, &Item{node: e.to, distance: nd})
			}
		}
	}
	// 4. return items
	return dist
}
