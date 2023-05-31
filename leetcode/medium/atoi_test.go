/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/2/11
 * @contact frankstarye@tencent.com
 **/

package medium

import (
	"fmt"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				s : "-----4567sssss",
			},
			want: 0,
		},
		{
			name: "test positive",
			args: args{
				s : "44 sss",
			},
			want: 44,
		},
		{
			name: "test 3",
			args: args{
				s : "   -42",
			},
			want: -42,
		},
		{
			name:"words and 987",
			args: args{
				s :"words and 987",
			},
			want: 0,
		},
		{
			name: "test 1",
			args: args{
				s : "+1",
			},
			want: 1,
		},
		{
			name: "testsss",
			args: args{
				s : "-+12",
			},
			want: 0,
		},
		{
			name: "\"  -0012a42\"",
			args: args{
				s :"  -0012a42",
			},
			want: -12,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got int
			if got = myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}
