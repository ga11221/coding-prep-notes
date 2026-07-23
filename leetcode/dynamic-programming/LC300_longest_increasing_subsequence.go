package main

import "fmt"

/*
i=0 0++LIS => [] => [0]
i=1 1++LIS where LIS[last] < 1 => [0, 1]
i=2 0++LIS where LIS[last] < 0 => [] =>[0]
i=3 3++LIS where LIS[last] < 3 => [0,1,3]
i=4 2++LIS where LIS[last] < 2
i=5 3++LIS where LIS[last] < 3

LIS[last] -> length(LIS)
i = 0 0 -> 1
i = 1 1 -> 2
i = 2 skip 0
i = 3 3 -> 3
i = 4 2 -> 3
i = 5 3 -> 4

[0, 1, 0, 3, 2, 3]

	 0  0  0  0  0  0
	    1     1  1  1
		      3  2  2
					3
*/
func lengthOfLIS(nums []int) int {
	LIS := map[int]int{}
	for _, num := range nums {
		lisLength := 0
		for last, length := range LIS {
			if last < num && length > lisLength {
				lisLength = length
			}
		}
		LIS[num] = lisLength + 1
	}
	lisLength := 0
	for _, length := range LIS {
		if length > lisLength {
			lisLength = length
		}
	}
	return lisLength
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	//nums := []int{0, 1, 0, 3, 2, 3}

	fmt.Print("longest subseq for: %v, has length: %d", nums, lengthOfLIS(nums))
}

/*
	    #300. Longest Increasing Subsequence

		Input: nums = [10,9,2,5,3,7,101,18]
		Input: nums = [10,9,2,10,3,7,6,18]
	       state - track longest subsequence start index and last number seen
	       init - state = [[10,1],[9,1], [2,1],[5,1], [3,1],[7,1],[101, 1], [18,1]] - Longest subsequence at every index is 1 with the last number in the subsequence being nums[i]
			or [][]state
		at i = 1, if nums[i] > state[i-1][0] state[i-1][1]++
		...
		at i = 3, state[3][1]++
	       in second tuple, store numbers encountered in order smaller than this num[i]
	       state=[[10,1],[9,1], [2,1],[4,[2]], [5,[2,4]], [3,[2]],[7,[2,3]],[101, [2,3,7]], [18,1]]

	       state=[[10,1],[9,1], [2,1],[10,[2]], [3,[]],[7,[3]],[6, [2,3]], [18,1]]
*/

/*
type state struct {
	num                          int
	longestSubSequenceEndingHere []int
}

func lengthOfLIS(nums []int) int {
	s := []*state{}
	for _, num := range nums {
		s = append(s, &state{
			num,
			[]int{},
		})
	}
	for i := 1; i < len(s); i++ {
		curr := s[i]
		last := s[i-1]
		if curr.num > last.num {
			curr.longestSubSequenceEndingHere = append(last.longestSubSequenceEndingHere, last.num) // modifies last and copies
		} else {
			lastSeq := last.longestSubSequenceEndingHere
			if len(lastSeq) > 0 {
				if lastSeq[len(lastSeq)-1] < curr.num {
					curr.longestSubSequenceEndingHere = append(curr.longestSubSequenceEndingHere, last.longestSubSequenceEndingHere...)
				}
			}
		}
	}
	max := 0
	for _, st := range s {
		length := len(st.longestSubSequenceEndingHere) + 1
		if length > max {
			max = length
		}
	}
	return max
}
*/
