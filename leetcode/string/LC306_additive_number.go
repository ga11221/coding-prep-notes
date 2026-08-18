package main

import "fmt"

/*
LC306 Additive Number (String, Backtracking)

An additive number is a string whose digits can form an additive sequence.

A valid additive sequence should contain at least three numbers. Except for
the first two numbers, every subsequent number in the sequence must be the
sum of the preceding two.

Given a string num consisting only of digits, return true if num is an
additive number. Otherwise, return false.

Input: num = "112358"
Output: true
Explanation: The digits form an additive sequence: 1, 1, 2, 3, 5, 8.
1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8.

Input: num = "199100199"
Output: true
Explanation: The digits form an additive sequence: 1, 99, 100, 199.
1 + 99 = 100, 99 + 100 = 199.

Constraints:
- 1 <= num.length <= 35
- num consists of only digits.
*/

/*

1. all permutations of first two pointers to mark first two numbers, permutations of the next pointer (within bounds) for the third and so on

1 1 2 3 5 8
11 2 3 5 8
11 23 5 8
11 235 8
11 2358
112 3 5 8 etc

2. if the members of the set are the indices of input, then semiring is (combine=add next to prev, extend=split next number)
3. recast/merge/drop - instead of trying every perm/index, local state/transition? has to be recast if avoiding permutations
   n^2 for first two numbers, then n for each subsequent number, so O(n^3) time complexity - with pruning for leading zeros and length of next number being too long, can be reduced to O(n^2) time complexity

*/

func main() {
	fmt.Println(isAdditiveNumber("112358"))    // true
	fmt.Println(isAdditiveNumber("199100199")) // true
	fmt.Println(isAdditiveNumber("1023"))      // false
}

func isAdditiveNumber(num string) bool {
	// for every possible first number, for every possible second number, check if the rest of the string can be formed by adding the two numbers
	// after selecting the first two numbers, we can check if the rest of the string is long enough to form a valid sum by adding the two numbers
	for i := 1; i <= len(num)/2; i++ {
		for j := i + 1; j < len(num); j++ {
			if isValid(num, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func isValid(num string, start int, i int, j int) bool {
	// check if the first two numbers are valid (no leading zeros)
	if (num[start] == '0' && i-start > 1) || (num[i] == '0' && j-i > 1) {
		return false
	}
	num1 := num[start:i]
	num2 := num[i:j]
	sum := addStrings(num1, num2)
	if len(sum) > len(num)-j || num[j:j+len(sum)] != sum {
		return false
	}
	if j+len(sum) == len(num) {
		return true
	}
	return isValid(num, i, j, j+len(sum))
}

func addStrings(num1 string, num2 string) string {
	// add two strings representing numbers
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	result := ""
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}
		carry = sum / 10
		result = string(sum%10+'0') + result
	}
	return result
}
