/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/8
 * @contact frankstarye@tencent.com
 **/

package arrays

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				x:121,
			},
			want: true,
		},
		{
			name: "test 2",
			args: args{
				x : 10,
			},
			want: false,
		},
		{
			name: "test 3",
			args: args{
				x :1001,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
