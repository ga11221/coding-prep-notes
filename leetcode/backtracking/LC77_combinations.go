package main

func combine(n int, k int) [][]int {
	combos := [][]int{}
	for i := 1; i <= n-k+1; i++ {
		_combine(&combos, []int{i}, i+1, n, k-1)
	}
	return combos
}

// might be easiest to frame recursive solutions as trees, descending one level is a function call
//      1			root = i = 1
//  2		3       L2 = i(2)..k(4)-n(2)+1
// 3 4			4	L3 = i(3)..k(4)-n(1)+1

// 2			root = i+1=2
//
//	  3			L2 = i+1..k-1
//		4		L3 = i+1..k
func _combine(combos *[][]int, arr []int, i int, n int, k int) {
	if k == 0 {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*combos = append(*combos, tmp)
		return
	}
	for ; i <= n-k+1; i++ {
		_combine(combos, append(arr, i), i+1, n, k-1)
	}
}
