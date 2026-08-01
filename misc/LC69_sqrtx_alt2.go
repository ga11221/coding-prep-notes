package main

import "testing"

func TestMySqrt(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{8, 2},
		{4, 2},
		{9, 3},
		{0, 0},
		{1, 1},
		{101, 10},
		{16, 4},
		{2, 1},
	}
	for _, tt := range tests {
		got := mySqrt(tt.input)
		if got != tt.want {
			t.Errorf("mySqrt(%d) = %d, want %d", tt.input, got, tt.want)
		}
	}
}
