package tencent

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "TestTwoSum: [2, 7, 11, 15], 9",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},

		{
			name: "TestTwoSum: [3, 2, 4], 6",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},

		{
			name: "TestTwoSum: [3,3], 6",
			args: args{
				nums:   []int{3,3},
				target: 6,
			},
			want: []int{0, 1},
		},

		{
			name: "TestTwoSum: [-3,4,3,90], 0",
			args: args{
				nums:   []int{-3,4,3,90},
				target: 0,
			},
			want: []int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}