/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/8
 * @contact frankstarye@tencent.com
 **/

package arrays

import "testing"

func Test_isReverseNumber(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test for reverse string",
			args: args{
				number: 123321,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isReverseNumber(tt.args.number); got != tt.want {
				t.Errorf("isReverseNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
