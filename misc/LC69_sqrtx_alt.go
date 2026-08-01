package main

func mySqrt(x int) int {
	sqrt := x / 2
	for sqrt*sqrt > x {
		sqrt /= 2
	}
	left := sqrt
	for sqrt*sqrt < x {
		left = sqrt
		sqrt++
	}
	if sqrt*sqrt == x {
		return sqrt
	}
	return left
}
