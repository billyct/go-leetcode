package tencent

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestLongestPalindrome: babad",
			args: args{
				s: "babad",
			},
			want: "bab",
		},

		{
			name: "TestLongestPalindrome: cbbd",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},

		{
			name: "TestLongestPalindrome: a",
			args: args{
				s: "a",
			},
			want: "a",
		},

		{
			name: "TestLongestPalindrome: ac",
			args: args{
				s: "ac",
			},
			want: "a",
		},

		{
			name: "TestLongestPalindrome: bb",
			args: args{
				s: "bb",
			},
			want: "bb",
		},

		{
			name: "TestLongestPalindrome: ",
			args: args{
				s: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("LongestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}