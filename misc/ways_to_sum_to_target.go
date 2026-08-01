package main

/**

@TODO 7/12
No general recurrence: The if/else chain (num==1, i==num, i%num==0, fallback) is a set of special cases masquerading as a recurrence. It breaks on any input not matching those patterns.

i%num == 0→ 2 + dp[num][i-1]: No clear justification for the 2+. Where does the 2 come from?

Semantics undefined: Is this combinations? Permutations? Unlimited usage? At-most-once? The computed values suggest permutations (e.g., dp[3][4]=2 for {1,3} and {3,1}), but the structure doesn't match any standard DP for counting sums.

This would fail on most test inputs. Should be rewritten with a clear state definition and a standard recurrence.


@TODO 7/13
consider how to change algo if dupes allowed in nums (ie either dedupe input first or
 freq + constraint ??? look into this second option further)
*/

import (
	"fmt"
)

func main() {
	nums := []int{3, 5}
	target := 7
	fmt.Printf("for nums: %v number of ways to sum to %d: %d\n", nums, target, sumToTarget(nums, target))
	_sumToTarget(nums, []int{}, target)
	fmt.Printf("for nums: %v _number of ways to sum to %d: %d\n", nums, target, len(globalSeqs))
	ways := 0
	__sumToTarget(nums, target, &ways)
	fmt.Printf("for nums: %v __number of ways to sum to %d: %d\n", nums, target, ways)
}

/*
	for 0 < A < B < C
	f(A, 0) - number of distinct sequences (permutations) that use at least one A and sum to 0
	f(B, 0) - number of distinct sequences (permutations) that use at least one B and sum to 0
	f(C, 0) - number of distinct sequences (permutations) that use at least one C and sum to 0

	f(A, 1) = 1(?) + f(A,0) if A = 1
	f(A, 2) = 1(?) + f(A,1) if A = 1
	f(A, 3) = 1(?) + f(A,2) if A = 1

	f(A, 1) = 0 + f(A,0) if A > 1
	f(A, 2) = 1 + f(A,0) if A = 2
	f(A, 3) = f(A,2) + f(B, t-2) if A =2 and B=1

	globalSeqs := [][]int
	sum(nums=[1,3,5], seq=[], 6) {
		if target == 0 &&  seq not in globalSeqs:
			add seq to globalSeqs
		sum1 = 0
		if 1 <= target {
			sum1 = sum([1,3,5], [1], 5)
		}
		sum3 = 0
		if 3 <= target {
			sum3 = sum([1,3,5], [3], 3)
		}
		sum5 = 0
		if 5 <= target {
			sum5 = sum([1,3,5], [5], 1)
		}
		return sum1+sum2+sum5
	}

	sum0 = sum (
		f(a,0) = 0,
		f(b,0) = 0,
		f(c,0) = 0,
	)
	------ or -----------
	sum0 = f(a,0) + f(b,0) + f(c,0)

	sum1 = sum (
		f(a,1) = 1+sum0 if target - sum0 >= a,
		f(b,1) = 1+sum0 if target - sum0 >= b,
		f(c,1) = 1+sum0 if target - sum0 >= c
	)

*/
// nums = sorted array of distinct integers > 0
// full of bugs - only works for one particular input/target
func sumToTarget(nums []int, target int) int {
	dp := make([][]int, nums[len(nums)-1]+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	for i := 1; i <= target; i++ {
		for n := 0; n < len(nums); n++ {
			num := nums[n]
			if target < num {
				continue
			}
			if num == 1 {
				dp[num][i] = 1
				continue
			}
			if i == num {
				dp[num][i] = 1
				continue
			}
			if i%num == 0 {
				dp[num][i] = 2 + dp[num][i-1]
				continue
			}
			dp[num][i] = 1 + dp[num][i-1]
		}
		/*
			dp[a][t] = number of distinct sequences of (a,b,c) that sum to t with at least one a

			dp[1][0]=0; dp[1][1]=1; dp[1][2]=1; dp[1][3]=1; dp[1][4]=1; dp[1][5]=1; dp[1][6]=1
			dp[3][0]=0; dp[3][1]=0; dp[3][2]=0; dp[3][3]=1; dp[3][4]=2; dp[3][5]=3; dp[3][6]=5
			dp[5][0]=0; dp[5][1]=0; dp[5][2]=0; dp[5][3]=0; dp[5][4]=0; dp[5][5]=1; dp[5][6]=2
		*/
	}
	sum := 0
	for n := 0; n < len(nums); n++ {
		sum += dp[nums[n]][target]
	}
	return sum
}

// perms order matters (ie {1,3} distinct from {3,1}) -
// so pick any num from state - no idx needed to track direction of movement in state
/*
__sumToTarget works - only one bounded param that changes - target - this is state
f(target) =
           for each num in nums:
		      f(target-num) if num <= target
base case: dp[0] = 1 // incrementing *ways when target = 0, so there is one way to get 0 - with empty set
answer in dp[target]
infer from base case that state def is: dp[n] = number of ways to sum to n

    shouldn't add 1 to prior states - the recursive function only does counting when target = 0 - and the "picking" is done in the for-loop
	dp[1] =      for each num in nums:
				  1+dp[0] if num <= target
	dp[1] = 1
	dp[2] = for each num in nums:
			  1+dp[1] if num <= target
    dp[2] = 2 // WRONG
	...


	dp[1] = sum over all n <= t=1 dp[t-n]
	dp[1] = 1
	dp[2] = sum over all n <= t=2 dp[t-n] = dp[1] = 1
    dp[2] = 1 // CORRECT!!!
	dp[3] = sum over all n <= t=3 dp[t-n] = sum(dp[3-3], dp[3-1])=sum(dp[0]+dp[2]) = sum(1, 1) = 2
*/

/*
	IMPORTANT!!!!!!!

in the dp recurrence, dp[t] = sum over n<=t: dp[t-n] => but it's not clear where in the recursive function this summing is taking place, if at all. how is that dp recurrence derived?
+ Thought: 5.0s
Great question. The summing is implicit in the recursion — that's why it's hard to see.
Trace __sumToTarget([1,3], 4) manually. Instead of ways, have each call return its contribution:

	func f(nums []int, target int) int {
	    if target == 0 { return 1 }           // base
	    total := 0
	    for _, n := range nums {
	        if n <= target {
	            total += f(nums, target-n)    // sum contributions
	        }
	    }
	    return total
	}

Now the total += is explicit. The original __sumToTarget hides this sum inside the pointer — each branch calls __sumToTarget which eventually either hits target==0 and increments *ways, or recurses further. The for loop fans out into branches, and *ways aggregates all base-case hits across the entire tree. That aggregation is the sum.
So dp[t] = sum dp[t-n] is just a direct translation of:
total := 0

	for _, n := range nums {
	    if n <= t {
	        total += f(t-n)
	    }
	}

return total
which is exactly the return-valued version of the pointer-accumulating recursion. The only difference is dp caches f(t) so you don't recompute it.
▣  Plan · Big Pickle · 11.7s
the lesson: don't propagate state in the params of a recursive func
+ Thought: 2.6s
Exactly the insight. Parameters that carry history (seq, path, curr) are usually IO accumulators, not state — they don't shrink the state space, they just collect results. Real DP state is only the params that determine what subproblem remains. If you can't drop the accumulator and still express the recurrence, your state definition is wrong.
*/
func __sumToTarget(nums []int, target int, ways *int) {
	if target == 0 {
		*ways = *ways + 1
		return
	}
	for _, n := range nums { // target = 3
		if n <= target { // n = 1,3
			__sumToTarget(nums, target-n, ways)
		}
	}
	/* f(3, 0)
		      |
	   for n from 1..3:
	   n=1:	f(2,0) ------------------> n=3: f(0,0)
		      |    <-------                ways++
	   for n from 1..1:	    ^              ways=2
			f(1,0)			|
		      |             |
	   for n from 1..1:	    |
			f(0,0)          |
			ways++      -----
			ways=1
	*/
}

/*
globalSeqs := [][]int

	sum(nums=[1,3,5], seq=[], 6) {
		if target == 0 &&  seq not in globalSeqs:
			add seq to globalSeqs
		sum1 = 0
		if 1 <= target {
			sum1 = sum([1,3,5], [1], 5)
		}
		sum3 = 0
		if 3 <= target {
			sum3 = sum([1,3,5], [3], 3)
		}
		sum5 = 0
		if 5 <= target {
			sum5 = sum([1,3,5], [5], 1)
		}
		return sum1+sum2+sum5
	}
*/
// can't be converted to dp easily - seq is unbounded/used to dedupe global var globalSeqs
var globalSeqs = [][]int{}

func _sumToTarget(nums []int, seq []int, target int) {
	if target == 0 {
		hasSeq := false
		for _, s := range globalSeqs {
			if len(s) == len(seq) {
				hasSeq = true
				for i, num := range s {
					hasSeq = hasSeq && num == seq[i]
				}
				if hasSeq {
					break
				}
			}
		}
		if !hasSeq {
			globalSeqs = append(globalSeqs, seq)
		}
		return
	}
	for _, n := range nums {
		if n <= target {
			cp := make([]int, len(seq))
			copy(cp, seq)
			cp = append(cp, n)
			_sumToTarget(nums, cp, target-n)
		}
	}
}
