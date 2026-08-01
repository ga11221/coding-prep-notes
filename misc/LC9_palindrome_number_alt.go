package main

import (
	"fmt"
	"math"
)

/*
Given an integer x, return true if x is a palindrome, and false otherwise.
*/

func main12() {
	//var num = 121; place = 2; % pow(place) = 21 , % 10 = 1
	// num = num / 10 = 12 , % pow(place - 1) = 2
	// num = num / 10 = 1

	//var num = -121

	var num = 11

	fmt.Print(isPalindrome(num))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	var num = x
	var place int
	for num > 9 {
		num = num / 10
		place += 1
	}

	var digits = []int{}
	var k = 0
	for i := place; i > 0; i-- {
		num = x / int(math.Pow10(k))
		for j := i; j > 0; j-- {
			num %= int(math.Pow10(j))
		}
		k++
		digits = append([]int{num}, digits...)
	}
	digits = append([]int{x / int(math.Pow10(k))}, digits...)

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}

	return true
}
