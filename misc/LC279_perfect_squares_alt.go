package main

func main() {}

/*
Given an integer n, return the least number of perfect square numbers that sum to n.

A perfect square is an integer that is the square of an integer;
in other words, it is the product of some integer with itself.
For example, 1, 4, 9, and 16 are perfect squares while 3 and 11 are not.
*/
func numSquares(n int) int {
	ps := []int{}
	for i := 0; i < n; i++ {
		if i*i <= n {
			ps = append(ps, i*i)
		} else {
			break
		}
	}
	return _numSquares(ps, n)
}
/*
    1. find largest ps ps1 in map < n
	subtract from n and find largest ps in map < n - ps1
	if no ps found, add 1s -> store number of terms in global min
	repeat for second largest ps ps1' in map < n
	
	repeat loop until ps1 = 1 (1s must have already been used in previous iterations)

ps = 1,4,9,16,25,36
	f(1) = 1
	f(4) = 1
	f(5) = 1+f(4)
	f(6) = 2+f(4)
	f(7) = 3+f(4)
	f(8) = 8 is multiple of perfect square 4 (largest perfect square < 8) = 
			4+f(4)=>f(4)+f(4)=2
	f(9) = 1
	f(10) = 1 + f(9) = 2
	      = 9+1
	f(11) = 2 + f(9) = 3
	      = 9+1+1
	f(12) = f(lps<12) + f(3)
		  = 9+1+1+1 or 4+4+4
	f(19) = 16+1+1+1 or 9+9+1
    f(20) = 16+4 or 9+9+1+1 or ...

	f(35) = 16+16+1+1+1 -> 5 terms or 3*9+2*4 -> 5 terms
	      = 25 + 9 + 1
	f(36) = 6*6 -> 2 terms or 4x9 -> 4 terms or 16+16+4 -> 3 terms
	
	perfectSquares = [1..n/2]

	f(n) = if n in perfectSquares -> 1

*/
func _numSquares(perfectSquares []int, n int) int {
	
}
