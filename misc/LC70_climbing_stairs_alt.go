package main

import "fmt"

/*
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

*/

func mainasdjfkkk() {
	//var target = 9

	//var target = 6

	var target = 4

	fmt.Print(climbStairs(target))
}

func climbStairs(n int) int {
	m := map[int]int{}
	return _climbStairs(n, m)
}

func _climbStairs(n int, m map[int]int) int {
	if n == 0 || n == 1 {
		return 1
	}
	/*
		n = 4
			n = 3
				n = 2
					n = 1
				n = 0
			n = 1
		n = 2

	*/

	if _, ok := m[n]; !ok {
		stairs := _climbStairs(n-1, m)
		stairsBy2 := _climbStairs(n-2, m)
		m[n] = stairs + stairsBy2
	}

	return m[n]
}

/*

n = 1
[1]

n = 2
[1,1] -> [2,1,1], [1,2,1], [1,1,2]
[2] -> [2,2]

n=3
[1,1,1] -> 1,1,1,1
[2,1]  -> 1,2,1, 2,1,1
[1,2] -> 1,1,2, 1,2,1


n=4
[1,1,1,1] -> 1 x 5

[2,1,1] -> [1,2,1,1], [2,1,1,1]
[1,2,1] -> [1,1,2,1], [1,2
[1,1,2]
[2,2] -> [1,1,2], [2,1,1]

n = 5
[1,1,1,1,1]

[1,1,1,2]
[1,1,2,1]
[1,2,1,1]
[2,1,1,1]


[2,2,1] -> [1,1,2,1], [2,1,1,1]
[1,2,2] -> [1,1,1,2], [1,2,1,1]
[2,1,2] -> [1,1,1,2], [2,1,1,1]

n = 6
[2,2,2] ->

n = 7
[2,2,2,1]
[2,2,1,2]
[2,1,2,2]
[1,2,2,2]

n = 4
[1,1,1,1]


[2,2]
 [2,1,1]
 [1,1,2]



[1
 [1,2
   [1,2,1]
 [1,1
	[1,1,1
      [1,1,1,1]
    [1,1,2]
[2
	[2,1
		[2,1,1]

	[2,2]
*/
