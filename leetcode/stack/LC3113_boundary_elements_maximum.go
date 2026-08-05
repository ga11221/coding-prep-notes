package main

import "fmt"

func main() {
	nums := []int{1, 4, 3, 3, 2}
	fmt.Println(numberOfSubarrays(nums)) // 6

	nums2 := []int{3, 3, 3}
	fmt.Println(numberOfSubarrays(nums2)) // 6

	nums3 := []int{1}
	fmt.Println(numberOfSubarrays(nums3)) // 1
}

/*
LC3113 Count Subarrays Where First and Last Elements Equal the Maximum
(see ladder.md - the generalized 6-step ladder).

Problem: given an array of positive integers, count subarrays where the first
and last elements are equal to the largest element in the subarray.

Rung 1 - ENUMERATE (ground truth):
  all O(n^2) subarrays; check s[0] == s[last] == max(s). Correctness anchor only.

Rung 2 - NAME THE OBJECT:
  count of witness pairs (i, j), i < j, with nums[i] == nums[j] and
  max(nums[i..j]) == nums[i]. A COUNTING problem: must finish the pass and keep
  a per-value count (multiset), never a flag (set).

Rung 3 - COMPRESS THE SPACE (representation axis):
  dominance, not distance: (i, j) qualifies iff no STRICTLY greater value lies
  between them, so a greater element permanently kills every equal-pair that
  would span it. Representation: a NON-INCREASING monotonic stack of values
  still "visible" - push num; pop strictly-smaller tops (now dominated); keep
  equals (they can still pair). Same boundary-decides-dominance move as
  LC84 largest-rectangle / daily-temperatures.

Rung 4 - COLLAPSE TO A STATISTIC (aggregation axis):
  seen[v] = count of active occurrences of v still in the stack. For each
  element: seen[num]++; count += seen[num] = 1 (the single) + one valid
  subarray per prior active occurrence. Each pop decrements seen of the
  dominated value.

Rung 5 - FIX A COMPUTATION ORDER (time axis):
  one left-to-right pass; each element enters and leaves the stack once, so the
  pops amortize to O(n). No early exit: count problems finish the pass.

Rung 6 - PROVE AN INVARIANT:
  before processing nums[i], the stack is non-increasing and holds exactly the
  prior values with no strictly-greater element after them, and seen[v] is
  their count. So every active occurrence of num pairs with nums[i] into a
  valid subarray, and each valid subarray is counted exactly once, at its right
  endpoint.

Complexity: O(n) amortized time, O(n) space.
*/
func numberOfSubarrays(nums []int) int {
	stack := []int{}
	var count int
	seen := map[int]int{}
	for _, num := range nums {
		top := len(stack) - 1
		if top < 0 {
			stack = append(stack, num)
		} else if num > stack[top] {
			i := top
			for i >= 0 && stack[i] < num {
				if _, ok := seen[stack[i]]; ok {
					seen[stack[i]]--
				}
				i--
			}
			stack = append(stack[0:i+1], num)
		} else {
			stack = append(stack, num)
		}
		seen[num]++
		count += seen[num]
	}
	return count
}
