package tencent

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		A []int
		m int
		B []int
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test_merge: A = [1,2,3,0,0,0], m = 3, B = [2,5,6], n = 3",
			args: args{
				A: []int{1,2,3,0,0,0},
				m: 3,
				B: []int{2,5,6},
				n: 3,
			},
			want: []int{
				1,2,2,3,5,6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.A, tt.args.m, tt.args.B, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
