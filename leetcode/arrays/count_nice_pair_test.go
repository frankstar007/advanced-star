/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/17
 * @contact frankstarye@tencent.com
 **/

package arrays

import "testing"

func Test_rev(t *testing.T) {
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
			name: "test rev",
			args: args{
				x: 1200,
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rev(tt.args.x); got != tt.want {
				t.Errorf("rev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countNicePairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test for nice pair",
			args: args{nums: []int{42,11,1,97}},
			want: 2,
		},
		{
			name: "test 2",
			args: args{
				nums: []int{13,10,35,24,76},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNicePairs(tt.args.nums); got != tt.want {
				t.Errorf("countNicePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}