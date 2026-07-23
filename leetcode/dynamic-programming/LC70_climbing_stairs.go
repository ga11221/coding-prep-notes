package main

func main() {}

// state dp(n)
// dp[1] = 1
// dp[2] = 2
// dp[3] = dp[2]+dp[1]=3
// dp[4] = dp[3]+dp[2]=5
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
