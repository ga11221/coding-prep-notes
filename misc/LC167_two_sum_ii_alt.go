package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("for numbers: %v and target: %v, values are %v", nums, target, twoSum(nums, target))
}

/*
	Example 1:

	Input: numbers = [2,7,11,15], target = 9
	Output: [1,2]
	Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

	Example 2:

	Input: numbers = [2,3,4], target = 6
	Output: [1,3]
	Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

	Example 3:

	Input: numbers = [-1,0], target = -1
	Output: [1,2]
	Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].


	f(i, j) = (i, j)          	  if nums[i] + nums[j] == target
        	= f(i, j-1)           if nums[i] + nums[j] > target
        	= f(i+1, j)           if nums[i] + nums[j] < target
*/

func twoSum(numbers []int, target int) []int {
	i, j := _twoSum(numbers, 0, len(numbers)-1, target)
	return []int{i + 1, j + 1}

}

func _twoSum(numbers []int, i, j, target int) (a, b int) {
	if numbers[i]+numbers[j] == target {
		return i, j
	}
	if numbers[i]+numbers[j] > target {
		return _twoSum(numbers, i, j-1, target)
	}
	return _twoSum(numbers, i+1, j, target)
}
