/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/7
 * @contact frankstarye@tencent.com
 **/

package arrays

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		nums []int
		x    int
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
				nums: []int{1,1,4,2,3},
				x: 5,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums, tt.args.x); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minOperation(t *testing.T) {
	type args struct {
		nums []int
		x    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test stack 1",
			args: args{
				nums: []int{1,1,4,2,3},
				x : 5,
			},
			want: 2,
		},
		{
			name: "test stack 2",
			args: args{
				nums: []int{5,6,7,8,9},
				x : 4,
			},
			want: -1,
		},
		{
			name: "test stack 3",
			args: args{
				nums: []int{3,2,20,1,1,3},
				x: 10,
			},
			want: 5,
		},
		{
			name: "test stack 4",
			args: args{
				nums: []int{3,2,20,1,1,3},
				x:10,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperation(tt.args.nums, tt.args.x); got != tt.want {
				t.Errorf("minOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}