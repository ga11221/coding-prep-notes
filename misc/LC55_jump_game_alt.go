package main

func main() {
	println(canJump([]int{2, 0, 0}))
}

func canJump(nums []int) bool {
	if nums[0] == 0 {
		return len(nums) == 1
	}
	dp := [][]int{}
	// dp contains start and end indices (inclusive) of every nZS in nums
	// tabulate nonZeroSegments
	// INVARIANT i => points to first index in this non-zero segment
	// INVARIANT j => points to first zero after this non-zero segment or last elem
	// INVARIANT k => points to last zero after this non-zero segment or last elem
	// dp stores every {i, j-1} pair

	for i := 0; i < len(nums); {
		j := i
		for ; j < len(nums) && nums[j] != 0; j++ {

		}
		if i == 0 && j == len(nums)-1 {
			return true
		}
		maxJump := 0
		for m := i; m < j; m++ {
			if m+nums[m] > maxJump {
				maxJump = m + nums[m]
			}
		}
		dp = append(dp, []int{i, j - 1, maxJump})
		k := j
		for ; k < len(nums) && nums[k] == 0; k++ {

		}
		i = k
	}

	return _canJump(dp, len(nums)-1)
}

/*
dp contains all nonZeroSegments and the maxJump index
let reachable = next reachable nZS(s)
if dp[0].maxJump >= last index return true
find all nZSs from 1..len(dp) where dp[0].maxJump >= nZS.end || dp[0].maxJump >= nZS.start

	for each nZSs in jumpable

if dp[0] > end return true
else find all dp[1:] that dp[0] is greater than (jumpable)

for each nZS in jumpable:

	if nZS > end
		return true
	else find all dp[nZS+1:] that nZS is greater than

	(dp[0], dp[1:], end)
		   |                       \
		   (dp[1], dp[2:], end)      (dp[2], dp[3:], end)
*/
func _canJump(dp [][]int, end int) bool {
	i := 0 // index of the next reachable segment
	reachable := [][]int{dp[i]}
	// INVARIANT - reachable contains next jumpable nZS(s)
	for len(reachable) > 0 {
		next := reachable[0]
		reachable = reachable[1:]
		if next[2] >= end {
			return true
		}
		// INVARIANT - j points to next nZS after this reachable nZS
		for j := i + 1; j < len(dp); j++ {
			if next[2] >= dp[j][0] || next[2] >= dp[j][1] {
				reachable = append(reachable, dp[j])
				i = j
			} else {
				break
			}
		}
	}
	// POST-INVARIANT = reachable is empty (ie there are no reachable nZSs left and the nZSs encountered don't go beyond end
	return false
}
