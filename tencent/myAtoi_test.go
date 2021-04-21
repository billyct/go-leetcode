package tencent

import (
	"testing"
)

func TestMyAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestMyAtoi: 42",
			args: args{
				str: "42",
			},
			want: 42,
		},

		{
			name: "TestMyAtoi:        -42",
			args: args{
				str: "       -42",
			},
			want: -42,
		},

		{
			name: "TestMyAtoi: 42",
			args: args{
				str: "42",
			},
			want: 42,
		},

		{
			name: "TestMyAtoi: 4193 with words",
			args: args{
				str: "4193 with words",
			},
			want: 4193,
		},

		{
			name: "TestMyAtoi: words and 987",
			args: args{
				str: "words and 987",
			},
			want: 0,
		},

		{
			name: "TestMyAtoi: -91283472332",
			args: args{
				str: "-91283472332",
			},
			want: -2147483648,
		},
		{
			name: "TestMyAtoi: +1",
			args: args{
				str: "+1",
			},
			want: 1,
		},
		{
			name: "TestMyAtoi: -+1",
			args: args{
				str: "-+1",
			},
			want: 0,
		},
		{
			name: "TestMyAtoi: 2147483647",
			args: args{
				str: "2147483647",
			},
			want: 2147483647,
		},
		{
			name: "TestMyAtoi: 20000000000000000000",
			args: args{
				str: "20000000000000000000",
			},
			want: 2147483647,
		},
		{
			name: "TestMyAtoi: 0000000000012345678",
			args: args{
				str: "0000000000012345678",
			},
			want: 12345678,
		},

		{
			name: "TestMyAtoi: ",
			args: args{
				str: "",
			},
			want: 0,
		},

		{
			name: "TestMyAtoi:    ",
			args: args{
				str: "        ",
			},
			want: 0,
		},

		{
			name: "TestMyAtoi: +",
			args: args{
				str: "+",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MyAtoi(tt.args.str); got != tt.want {
				t.Errorf("MyAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}