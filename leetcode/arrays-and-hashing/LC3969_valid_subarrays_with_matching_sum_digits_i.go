package arrays

func countValidSubarrays(nums []int, x int) int {
	count := 0
	var sum int64
	for i := 0; i < len(nums); i++ {
		sum = int64(nums[i])
		if firstDigit(sum) == x && lastDigit(sum) == x {
			count++
		}
		for j := i + 1; j < len(nums); j++ {
			sum += int64(nums[j])
			if firstDigit(sum) == x && lastDigit(sum) == x {
				count++
			}
		}
	}
	return count
}

func firstDigit(num int64) int {
	for num >= 10 {
		num /= 10
	}
	return int(num)
}

func lastDigit(num int64) int {
	return int(num % 10)
}
