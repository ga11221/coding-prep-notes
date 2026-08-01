package main

func main() {}

// state dp(n)
// dp[0] = 0
// dp[1] = 1
// dp[2] = 1
// dp[3] = dp[2]+dp[1]+dp[0]
func tribonacci(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1] + dp[i-3]
	}
	return dp[n]
}
