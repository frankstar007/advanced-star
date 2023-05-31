/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/2/10
 * @contact frankstarye@tencent.com

	case1 : PAYPALISHIRING

			P     A     H     N
			A  P  L  S  I  I  G
			Y     I     R
 **/

package medium

import (
	"fmt"
	"testing"
)

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				s : "PAYPALISHIRING",
				numRows: 3,
			},
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "test 2",
			args: args{
				s : "PAYPALISHIRING",
				numRows: 4,
			},
			want: "PINALSIGYAHRPI",
		},
		{
			name: "test 3",
			args: args{
				s: "A",
				numRows: 1,
			},
			want: "A",
		},
		{
			name: "test 4",
			args: args{
				s: "ABCDE",
				numRows: 4,
			},
			want: "ABCED",
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if got = convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
			fmt.Println("base: " + tt.args.s)
			fmt.Println()
			fmt.Println("want: " + tt.want)
			fmt.Println()
			fmt.Println("mine: " + got)
		})
	}
}
