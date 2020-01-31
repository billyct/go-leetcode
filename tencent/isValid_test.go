package tencent

import "testing"

func TestIsValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test for ()",
			args: args{s: "()"},
			want: true,
		},

		{
			name: "test for ()[]{}",
			args: args{s: "()[]{}"},
			want: true,
		},

		{
			name: "test for (]",
			args: args{s: "(]"},
			want: false,
		},

		{
			name: "test for ([)]",
			args: args{s: "([)]"},
			want: false,
		},

		{
			name: "test for {[]}",
			args: args{s: "{[]}"},
			want: true,
		},

		{
			name: "test for ((",
			args: args{s: "(("},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.s); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}