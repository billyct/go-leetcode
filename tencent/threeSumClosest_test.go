package tencent

import (
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestThreeSumClosest: [-1, 2, 1, -4], 1",
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			want: 2,
		},

		{
			name: "TestThreeSumClosest: [0, 2, 1, -3], 1",
			args: args{
				nums:   []int{0, 2, 1, -3},
				target: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThreeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("ThreeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}