/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/5/25
 * @contact frankstarye@tencent.com
 **/

package easy

import "testing"

func Test_oddString(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test for odd",
			args: args{
				words: []string{"adc", "wzy", "abc"},
			},
			want: "abc",
		},
		{
			name: "test for sec",
			args: args{
				words: []string{"aaa", "bob", "ccc", "ddd"},
			},
			want: "bob",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddString(tt.args.words); got != tt.want {
				t.Errorf("oddString() = %v, want %v", got, tt.want)
			}
		})
	}
}
