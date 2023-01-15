/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/10
 * @contact frankstarye@tencent.com
 **/

package sort

import "testing"

func Test_heapSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		//TODO: Add test cases.
		{
			name: "heap sort",
			args: args{
				nums: []int{0,0,0},
			},
		},
		{
			name: "quick sort",
			args: args{
				nums: []int{-1,0,1,2,-1,-4},
			},
		},
		{
			name: "sss",
			args: args{
				nums: []int{1,1,1},
			},
		},
	}
	for _, tt := range tests {
		threeSum(tt.args.nums)
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
