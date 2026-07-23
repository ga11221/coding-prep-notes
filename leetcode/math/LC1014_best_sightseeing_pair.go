package main

// values = [8,1,5,2,6]
//          [8,0,3,-1,2]=>sort=> [8]++   [3,2,0,-1] max = 11
//                                     i= 2,4,1, 3
//                                    map original idx to where it occurs in sorted list
// or don't sort and bin search on [0,3,-1,2]
//          [8,1,4,0,3]       => [8,1]++  [4,3,0] max = 5
//                                      i= 2,4,3
//          [8,1,5,1,4]       => [8,1,5]++{4,1] max = 9
//                                      i= 4,3

// looking for all pairwise combos (i, j) where j > i
// prune or eliminate any?

// if i=0 and j=1
// sum = 8+1+0 -1 = 8
// value at j=2 must be at least 3 to be > 8
// sum = 8+3+0-2=9


// So given an array of integers, find the two that produce the max value 
// for values[i] + values[j] + i - j
// so for eg if i=2 and j=4 then values[2]+values[4]+2-4= 5+6-2 = 9
// find the i and j that produce the max
// elements can be negative? shouldn't affect procedure
// i != j

// Edge cases: empty list, singleton list, list of length 2
func maxScoreSightseeingPair(values []int) int {
	// [8, 1, 5]
	// maximize sum of elems while minimizing distance
	// try every possibility:
	// i=0,j=1
	// 8+1-1 =8
	// 8+5-2=11
	// 1+5-1=5
	// O(n^2)
	// 1. for each number except the last in values:
	//  2. sum with nums to its righth and subtract i-j from that sum
	//  3. if sum from 2. is larger then global max, set global max to sum

	// edge cases: values guaranteed to be of length 2 or greater
	// values are positive and > 0, can init global max to 0
	globalMax := 0
	for i:= 0; i < len(values) -1;i++ {
		for j:=i+1;j<len(values);j++ {
			sum := values[i]+values[j]+(i-j)
			if sum > globalMax {
				globalMax = sum
			}
		}
	}
	return globalMax
	
}
