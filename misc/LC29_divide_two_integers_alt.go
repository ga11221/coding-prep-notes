package main

import (
	"fmt"
	"math"
)

func mainoaisdjij() {
	var dividend = 10
	var divisor = 3
	element := divide(dividend, divisor)
	fmt.Print(element)
	fmt.Print("\n")
}

func divide(dividend int, divisor int) int {
	if divisor == 1 {
		if dividend < 0 {
			return int(math.Max(float64(dividend), math.MinInt32))
		}
		return int(math.Min(float64(dividend), math.MaxInt32))
	}
	if divisor == -1 {
		if dividend < 0 {
			return int(math.Min(float64(0-dividend), math.MaxInt32))
		}
		return int(math.Max(float64(0-dividend), math.MinInt32))
	}

	var diff = int(math.Abs(float64(dividend)))
	var absDiv = int(math.Abs(float64(divisor)))
	i := 0
	for ; diff >= absDiv; i++ {
		diff -= absDiv
	}
	// both neg
	if (dividend < 0 && divisor < 0) || (dividend > 0 && divisor > 0) {
		return int(math.Min(float64(i), math.MaxInt32))
	}
	return int(math.Max(float64(0-i), math.MinInt32))
}
