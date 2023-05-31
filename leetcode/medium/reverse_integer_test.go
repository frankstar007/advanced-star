/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/2/10
 * @contact frankstarye@tencent.com
 **/

package medium

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				x : 123,
			},
			want: 321,
		},
		{
			name: "test 2",
			args: args{
				x : -123,
			},
			want: -321,
		},
		{
			name: "test 3",
			args: args{
				x : 120,
			},
			want: 21,
		},
		{
			name: "test s",
			args: args{
				x : -2147483648,
			},
			want: 0,
		},
		{
			name: "test e",
			args: args{
				x : 1463847412,
			},
			want: 2147483641,
		},
		{
			name: "test a",
			args: args{
				x : 1463847412,
			},
			want: 2147483641,

		},
		{
			name: "test s",
			args: args{
				x : 1534236469,
			},
			want: 0,
		},



	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
