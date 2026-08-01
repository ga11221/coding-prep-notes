package main

import (
	"testing"
)

func TestStockSpanner(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   []int
	}{
		name:   "100,80,60,70,60,75,85",
		prices: []int{100, 80, 60, 70, 60, 75, 85},
		// sort with pos: [(100,1)]
		//  [(80,2), (100,1)]
		//  [(60,3), (80,2), (100,1)]
		//  [(60,3), (70,4),(80,2), (100,1)]
		//  [(60,3),(60,5),(70,4), (80,2), (100,1)]
		//prices: []int{81, 80, 60, 70, 60, 75, 85},
		//			1,  1,  1,  2(stop at 60 b/c==1), 1,
		// prices: []int{100,99,98,97,96,95,94}
		// 98, 99, 100
		// 2,  1,  0
		//                1   1  1  1  1  1  1
		want: []int{1, 1, 1, 2, 1, 4, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := Constructor()
			for i, v := range prices {
				param_1 := obj.next(v)
				if param_1 != tt.want[i] {
					t.Errorf("stockSpanner(%vith) = %v, want %v", i, param_1, tt.want)
					return
				}
			}
			//obj.next(100) // return 1
			//obj.next(80)  // return 1
			//obj.next(60)  // return 1
			//obj.next(70)  // return 2
			//obj.next(60)  // return 1
			//obj.next(75)  // return 4, because the last 4 prices (including today's price of 75) were less than or equal to today's price.
			//obj.next(85)  // return 6
		})
	}
}
