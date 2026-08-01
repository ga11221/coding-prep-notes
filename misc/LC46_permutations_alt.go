package main

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	var res [][]int
	for i, n := range nums {
		rest := append(append([]int{}, nums[:i]...), nums[i+1:]...)
		for _, p := range permute(rest) {
			res = append(res, append([]int{n}, p...))
		}
	}
	return res
}
