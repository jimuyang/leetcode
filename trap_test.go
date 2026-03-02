package main

import (
	"testing"
)

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "空数组",
			height: []int{},
			want:   0,
		},
		{
			name:   "单个元素",
			height: []int{1},
			want:   0,
		},
		{
			name:   "LeetCode 示例1",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			name:   "LeetCode 示例2",
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
		{
			name:   "全零",
			height: []int{0, 0, 0},
			want:   0,
		},
		{
			name:   "单调递增-无积水",
			height: []int{1, 2, 3, 4, 5},
			want:   0,
		},
		{
			name:   "单调递减-无积水",
			height: []int{5, 4, 3, 2, 1},
			want:   0,
		},
		{
			name:   "单峰",
			height: []int{3, 0, 2, 0, 4},
			want:   7,
		},
		{
			name:   "两个等高柱",
			height: []int{2, 0, 2},
			want:   2,
		},
		{
			name:   "全部等高",
			height: []int{2, 2, 2, 2},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trap(tt.height)
			if got != tt.want {
				t.Errorf("trap(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}
