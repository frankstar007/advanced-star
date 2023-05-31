/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/2/20
 * @contact frankstarye@tencent.com
 **/

package medium

import "testing"

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
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
				n : 5,
			},
			want: "111221",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay(tt.args.n); got != tt.want {
				t.Errorf("countAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
