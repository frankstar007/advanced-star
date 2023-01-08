/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/4
 * @contact frankstarye@tencent.com
 **/

package binary_search

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "zero",
			args: args{
				nums: []int{1,3,5,6},
				target: 0,
			},
			want: 0,
		},
		{
			name: "test for search insert position",
			args: args{
				nums: []int{1,3,5,6},
				target: 7,
			},
			want: 4,
		},
		{
			name: "sec",
			args: args{
				nums: []int{1,3,5,6},
				target: 2,
			},
			want: 1,
		},
		{
			name: "third",
			args: args{
				nums: []int{1,3,5,6},
				target: 5,
			},
			want: 2,
		},
		{
			name: "fourth",
			args: args{
				nums: []int{1},
				target: 1,
			},
			want: 0,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
