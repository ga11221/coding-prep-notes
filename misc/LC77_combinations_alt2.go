package main

import (
	"slices"
	"testing"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want [][]int
	}{
		{
			name: "n=2, k=4",
			n:    4,
			k:    2,
			want: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		}, {
			name: "n=1, k=1",
			n:    1,
			k:    1,
			want: [][]int{{1}},
		}, {
			name: "n=3, k=4",
			n:    4,
			k:    3,
			want: [][]int{
				{1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combine(tt.n, tt.k)
			t.Logf("combine(%v, %v)=%v", tt.n, tt.k, got)
			if len(got) != len(tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
				return
			}
			for i := 0; i < len(got); i++ {
				if !slices.Equal(got[i], tt.want[i]) {
					t.Errorf("combine() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
