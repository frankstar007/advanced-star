/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/11/22
 * @contact frankstarye@tencent.com
 **/

package arrays

import "testing"

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "快速排序",
			args: args{
				nums: []int{3,2,1,5,6,4},
				k: 2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
