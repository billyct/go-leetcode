package tencent

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestLongestCommonPrefix: [\"flower\",\"flow\",\"flight\"]",
			args: args{
				ss: []string{"flower","flow","flight"},
			},
			want: "fl",
		},

		{
			name: "TestLongestCommonPrefix: [\"dog\",\"racecar\",\"car\"]",
			args: args{
				ss: []string{"dog","racecar","car"},
			},
			want: "",
		},

		{
			name: "TestLongestCommonPrefix: [\"aa\", \"a\"]",
			args: args{
				ss: []string{"aa", "a"},
			},
			want: "a",
		},

		{
			name: "TestLongestCommonPrefix: [\"a\"]",
			args: args{
				ss: []string{"a"},
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonPrefix(tt.args.ss); got != tt.want {
				t.Errorf("LongestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}