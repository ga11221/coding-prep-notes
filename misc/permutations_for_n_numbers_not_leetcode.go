package main

import "fmt"

func main() {
	fmt.Println(perms(1))
	println()
}

// [1,2,3]
// [1]++perms[2,3]
// [1]++[2]++perms[3]
// place 2 in all 2 positions of [3]
// [1]++[2,3] and [1]++[3,2]
// place 1 in all three positions of [2,3] and [3,2]
// [1]++[2,3] and [2,3]++[1] and [1]++[3,2] and [3,2]++[1]
// [1,2,3],[2,1,3], [2,3,1], [1,3,2],[3,1,2], [3,2,1]
func perms(n int) [][]int {
	nums := make([]int, n)
	allPerms := [][]int{}
	for i := range n {
		nums[i] = i
	}

	return nil
}
