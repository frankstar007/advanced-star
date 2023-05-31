/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/2/3
 * @contact frankstarye@tencent.com
 **/

package medium

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test for longest palindrome",
			args: args{
				s : "babad",
			},
			want: "bab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindromic(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test for 2",
			args: args{
				s:"abcdssssdcba",
			},
			want: "abcdssssdcba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromic(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromic() = %v, want %v", got, tt.want)
			}
		})
	}
}