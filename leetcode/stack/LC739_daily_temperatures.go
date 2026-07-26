package main

import "fmt"

func main() {
	temps := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temps)) // [1,1,4,2,1,1,0,0]

	temps2 := []int{30, 40, 50, 60}
	fmt.Println(dailyTemperatures(temps2)) // [1,1,1,0]

	temps3 := []int{30, 60, 90}
	fmt.Println(dailyTemperatures(temps3)) // [1,1,0]
}

/*
Given an array of integers temperatures representing the daily temperatures,
return an array answer such that answer[i] is the number of days you have
to wait after the ith day to get a warmer temperature. If there is no future
day for which this is possible, keep answer[i] == 0 instead.
*/

func dailyTemperatures(temperatures []int) []int {
	stack := []int{0}
	days := make([]int, len(temperatures))
	for i := 1; i < len(temperatures); i++ {
		temp := temperatures[i]
		top := len(stack) - 1
		for top >= 0 && temperatures[stack[top]] < temp {
			days[stack[top]] = i - stack[top]
			stack = stack[:top]
			top--
		}
		stack = append(stack, i)
	}
	for i := 0; i < len(stack); i++ {
		days[stack[i]] = 0
	}
	return days
}
