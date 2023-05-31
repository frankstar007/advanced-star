/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/2/12
 * @contact frankstarye@tencent.com
 **/

package medium

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				num: 1994,
			},
			want: "MCMXCIV",
		},
		{
			name: "test2",
			args: args{
				num: 58,
			},
			want: "LVIII",
		},
		{
			name: "test3",
			args: args{
				num: 3,
			},
			want: "III",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
