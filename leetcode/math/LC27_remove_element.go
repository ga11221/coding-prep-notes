package main

import "fmt"

func main32j() {
	var nums = []int{0, 1, 2, 2, 3, 0, 4, 2} // 2, 2,  ||||    2, 2, 2, 3
	var val = 2

	element := removeElement(nums, val)
	fmt.Print(element)
	fmt.Print("\n")
}

func removeElement(nums []int, val int) int {
	var numsNotMatching = []int{}
	for _, num := range nums {
		if num != val {
			numsNotMatching = append(numsNotMatching, num)
		}
	}
	for i := 0; i < len(numsNotMatching); i++ {
		nums[i] = numsNotMatching[i]
	}
	return len(numsNotMatching)
}
