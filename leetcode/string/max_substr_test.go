/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/11/20
 * @contact frankstarye@tencent.com
 **/

package string

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test first case",
			args: args{
				s : "abcabcbb",
			},
			want: 3,
		},

		{
			name: "test sec case",
			args: args{
				s : "bbbbb",
			},
			want: 1,
		},
		{
			name: "test third case",
			args: args{
				s : "pwwkew",
			},
			want: 3,
		},
		// TODO: Add test cases.
		{
			name: "test for longest sub string",
			args: args{
				s : "abssfdfdesekf",
			},
			want: 4,
		},

		{
			name: "test for longest sub string",
			args: args{
				s : "dvdf",
			},
			want: 3,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
