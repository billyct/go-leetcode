package tencent

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "TestThreeSum: [-1, 0, 1, 2, -1, -4]",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},

		{
			name: "TestThreeSum: [0,0,0]",
			args: args{
				nums: []int{0,0,0},
			},
			want: [][]int{
				{0,0,0},
			},
		},

		{
			name: "TestThreeSum: [1, -1, -1, 0]",
			args: args{
				nums: []int{1, -1, -1, 0},
			},
			want: [][]int{
				{-1, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThreeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThreeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}