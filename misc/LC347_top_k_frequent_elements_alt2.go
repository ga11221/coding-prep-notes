package main

import (
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "[1,2,1,2,1,2,3,1,3,2]",
			nums: []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "[3,0,1,0]",
			nums: []int{3, 0, 1, 0},
			k:    1,
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)
			slices.Sort(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedMap(t *testing.T) {
	gt := func(v1, v2 int) bool {
		return v1 >= v2
	}
	om := NewOrderedMap(gt, 2)
	om.put(1, 3)
	om.put(2, 2)
	om.put(3, 10)
	//t.Debug(om.OrderedKeys)

}
