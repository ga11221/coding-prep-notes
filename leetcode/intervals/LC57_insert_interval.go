package main

import "fmt"

/*
*
1. Type out the algo in comments
2. As edge cases arise (if they don't, look for them), create test cases
3. Write the code, walk interviewers through tests cases (edge cases first)
4. Space and time complexity
*/
func main() {
	newInterval := []int{0, 0}
	intervals := [][]int{{1, 5}}
	fmt.Print("intervals: %v\n", insert(intervals, newInterval))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	/**
	Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
	Output: [[1,5],[6,9]]
	Example 2:

	// 1. find index I where either:
	// 	a. newInterval.start fits inside an interval => merge
	// 	b. newInterval.start fits between a gap in intervals (ie newInterval.start < intervals[i].start) => insert
	// 2. from index I to end of intervals, search for index J where either:
	//  a. newInterval.end < some later interval.start => insert
	// 	b. newInterval.end fits inside an interval of intervals => merge
	/**
		1a.
		  [[1,3],[6,9]] nI = [2,5]
		     ^ i = 0 1a is true
			 	   ^ j starts at i+1 2b is true
				   merge I, insert J => keep interval[I].start, add interval[J].end and remove intervals from I to J-1, insert new interval in I
		  [[1,3], [6,9], [12,15]] nI = [2, 13]
		     ^ i = 0 1a is true
			        ^ j = 1 2a and 2b are false
					        ^ j = 2 2b is true
							merge I, merge J => keep interval[I].start, keep interval[J].end and remove intervals from I to J inclusive, insert newInterval in I

		i = next unexamined interval, boundary check for 1b

		Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
		Output: [[1,2],[3,10],[12,16]]
		Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
	*/
	indexForNewIntervalStart := -100 // I -1 sentinel collides with ... don't use magic number in interview
	indexForNewIntervalEnd := -100   // J b/c these start at -1, handle the empty case explicitly
	newMergingInterval := []int{}
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	for i := 0; i < len(intervals); i++ {
		interval := intervals[i]
		if newInterval[0] < interval[0] {
			indexForNewIntervalStart = i
			newMergingInterval = append(newMergingInterval, newInterval[0])
			break
		}
		if newInterval[0] < interval[1] {
			indexForNewIntervalStart = i
			newMergingInterval = append(newMergingInterval, interval[0])
			break
		}
	}
	if indexForNewIntervalStart == -100 { // didn't catch this edge case in design
		lastInterval := intervals[len(intervals)-1]
		if newInterval[0] == lastInterval[1] {
			lastInterval[1] = newInterval[1]
			return intervals
		}
		return append(intervals, newInterval)
	}
	for i := indexForNewIntervalStart; i < len(intervals); i++ {
		interval := intervals[i]
		if newInterval[1] < interval[0] {
			indexForNewIntervalEnd = i - 1
			newMergingInterval = append(newMergingInterval, newInterval[1])
			break
		}
		if newInterval[1] < interval[1] {
			indexForNewIntervalEnd = i
			newMergingInterval = append(newMergingInterval, interval[1])
			break
		}
	}
	if indexForNewIntervalEnd == -100 {
		indexForNewIntervalEnd = len(intervals) - 1
		newMergingInterval = append(newMergingInterval, newInterval[1])
	}
	for i := indexForNewIntervalStart; i <= indexForNewIntervalEnd; i++ {
		fmt.Print("interval start: %v\n", indexForNewIntervalStart)
		fmt.Print("interval end: %v\n", indexForNewIntervalEnd)
		intervals[i] = nil
	}
	if indexForNewIntervalStart == 0 && indexForNewIntervalEnd == -1 {
		return append([][]int{newMergingInterval}, intervals...)
	}
	intervals[indexForNewIntervalStart] = newMergingInterval
	newIntervals := [][]int{}
	for _, interval := range intervals {
		if interval != nil {
			newIntervals = append(newIntervals, interval)
		}
	}
	return newIntervals
}
