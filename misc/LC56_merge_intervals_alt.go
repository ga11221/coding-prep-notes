package main

import (
	"fmt"
	"sort"
)

func main() {
	// [[1,3],[2,6],[8,10],[15,18]]
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	// [[1,4],[4,5]]
	// intervals := [][]int{{1, 4}, {4, 5}}
	// [[4,7],[1,4]]
	// intervals := [][]int{{4, 7}, {1, 4}}
	fmt.Printf("merged intervals for %v: %v\n", intervals, merge(intervals))
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	// start with pointer1 at first interval, pointer2 at second interval
	// while there is an overlap between intervals[i] and intervals[j], merge and increment j
	// overlap defined as s1 < s2 and s2 < e1 and e1 < e2 (staggered) or s1 < s2 and e2 < e1 (enclosing)
	// if overlap found, if staggered, replace intervals[i] with (intervals[i].start, intervals[j].end) and replace intervals[j] with nil, if enclosing -> keep intervals[i] and replace intervals[j] with nil
	// if no overlap, increment i until i points to next non-nil interval and set j = i+1
	// ie i always points to first unmerged/unexamined interval, and j always resets to the next interval
	// length = 3
	// i = 0, j = 1
	// guard = 2
	for i, j := 0, 1; i < len(intervals)-1; {
		unmergedInterval := intervals[i]
		intervalToMerge := intervals[j]
		mergedInterval, overlaps := isOverlapping(unmergedInterval, intervalToMerge)
		if overlaps {
			intervals[i] = mergedInterval
			intervals[j] = nil
		}
		if j < len(intervals)-1 {
			j++
		} else {
			i++
			for ; i < len(intervals)-1 && intervals[i] == nil; i++ {
			}
			j = i + 1
		}
	}
	mergedIntervals := [][]int{}
	for _, interval := range intervals {
		if interval != nil {
			mergedIntervals = append(mergedIntervals, interval)
		}
	}
	return mergedIntervals
}

func isOverlapping(unmergedInterval, intervalToMerge []int) ([]int, bool) {
	if unmergedInterval[0] <= intervalToMerge[0] &&
		intervalToMerge[0] <= unmergedInterval[1] &&
		unmergedInterval[1] <= intervalToMerge[1] {
		return []int{unmergedInterval[0], intervalToMerge[1]}, true
	}
	if unmergedInterval[0] <= intervalToMerge[0] &&
		intervalToMerge[1] <= unmergedInterval[1] {
		return unmergedInterval, true
	}
	return nil, false
}

/**
// create table for lookup by window width
windowWidths := map[int][]int{}
for i, interval := range intervals {
	width := interval[1] - interval[0]
	if intervalIndices, ok := windowWidths[width] {
		intervalIndices = append(intervalIndices, i)
	} else {
		windowWidths[width] = []int{i}
	}
}
// merge enclosing intervals - ie tombstone enclosed intervals
widths := []int{}
for width := range windowWidths {
	widths = append(widths, width)
}
sort.Sort(sort.Reverse(sort.IntSlice(widths)))
// for every width at index i of widths, every width at index j > i is less than it
// 1. lookup the interval indices associated with the smaller widths
// 2. for every interval with width at index i, and every interval at smaller widths determine if (sj - sI) + v <= ej and replace with nil
for _, width := range widths  {


}
*/
