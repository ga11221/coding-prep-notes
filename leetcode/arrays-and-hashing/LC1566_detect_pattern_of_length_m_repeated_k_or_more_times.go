package main

/*
LC1566 Detect Pattern of Length M Repeated K or More Times (Array, Enumeration)

Given an array of positive integers arr, find a pattern of length m that is
repeated k or more times.

A pattern is a subarray (consecutive sub-sequence) that consists of one or more
values, repeated multiple times consecutively without overlapping. A pattern is
defined by its length and the number of repetitions.

Return true if there exists a pattern of length m that is repeated k or more
times, otherwise return false.

Input: arr = [1,2,4,4,4,4], m = 1, k = 3
Output: true
Explanation: The pattern of length 1 is 4, repeated 4 times consecutively.
[4,4,4,4] counts as k=3 or more repetitions of pattern [4].

Input: arr = [1,2,1,2,1,1,1,3], m = 2, k = 2
Output: true
Explanation: The pattern of length 2 is [1,2], repeated 2 times consecutively.

Input: arr = [1,2,1,2,1,3], m = 2, k = 3
Output: false
Explanation: The pattern of length 2 is [1,2], but it only occurs 2 times.
There is no pattern of length 2 repeated 3 times consecutively.

Input: arr = [1,2,3,1,2], m = 2, k = 2
Output: false
Explanation: Pattern [1,2] occurs 2 times but not consecutively; [3,1]
occurs once. No pattern of length 2 is repeated 2 times consecutively.

Input: arr = [2,2,2,2], m = 2, k = 3
Output: false
Explanation: The pattern of length 2 is [2,2], but it only occurs 2 times
consecutively. There is no pattern of length 2 repeated 3 times.

Constraints:
- 2 <= arr.length <= 100
- 1 <= arr[i] <= 100
- 1 <= m <= 100
- 2 <= k <= 100
*/

/*
Correct approach - fixed m*k window, linear scan with jump.

Key observation: a valid pattern has length exactly m*k (k copies of a
length-m block, non-overlapping, consecutive).

Approach:
  slide a fixed window of width m*k over arr; window [start, start+m*k) is a
  valid pattern iff arr[t] == arr[t+m] for every offset t in
  [start, start+m*(k-1)). First witness -> early return (existence problem).

  On a mismatch at index j, advance start to j+1 (NOT start+1): the pair
  (j, j+m) lies inside the comparison range of every window whose start s is
  in [start, j] (since s >= start >= j - m*k + m + 1), so all of those starts
  are provably invalid. This jump makes the scan O(n): each index is compared
  at most once as the left element of a pair.

Invariant: every pattern start < i has been proven invalid (a mismatch at
some j <= that start lies inside its comparison range). i never skips a valid
start.

6-rung mapping (see ladder.md):
  Rung 1 - ENUMERATE: all O(n) candidate starts s with s+m*k <= n, each
           checked over m*(k-1) offsets -> O(n*m*k). Ground truth.
  Rung 2 - NAME THE OBJECT: a window of width m*k that is k consecutive
           copies of a length-m block; question type EXIST (early return on
           first witness is a rung-2 property, not an optimization).
  Rung 3 - COMPRESS THE SPACE (representation): fixed width m*k collapses
           each candidate start to a boolean predicate; the mismatch jump
           (ordering) keeps the scan lossless and O(n).
  Rung 4 - COLLAPSE TO A STATISTIC (aggregation): the predicate is a boolean
           "window is k repetitions"; OR aggregation over starts.
  Rung 5 - FIX A COMPUTATION ORDER (time axis): single left-to-right pass;
           on mismatch at j, skip the whole range [i, j].
  Rung 6 - PROVE AN INVARIANT: "all starts < i are invalid" -> soundness
           (a passing window IS k repetitions) and completeness (no valid
           start is skipped by the jump).

Soundness: a window passing all arr[t] == arr[t+m] checks is exactly k copies
of the block [start, start+m).
Completeness: the mismatched pair (j, j+m) invalidates every start in [i, j],
so jumping to j+1 loses nothing.
Complexity: O(n) time, O(1) space.
*/

func main() {
	arr := []int{1, 2, 1, 2, 1, 3}
	m := 2
	k := 3
	containsPattern(arr, m, k)
}

func containsPattern(arr []int, m int, k int) bool {
	windowWidth := m * k
	if m*k > len(arr) {
		return false
	}
next:
	for i := 0; i+(windowWidth-1) < len(arr); {
		for j := i; j+m <= i+(windowWidth-1); {
			if j+m >= len(arr) {
				return false
			}
			if arr[j] != arr[j+m] {
				i = j + 1
				continue next
			}
			j++
		}
		return true
	}
	return false
}
