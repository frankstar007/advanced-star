/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/11
 * @contact frankstarye@tencent.com
 **/

package sort

import "testing"

func Test_quickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test for quick sort",
			args: args{
				nums: []int{30,40,60,10,20,50},
			},
		},
	}
	for _, tt := range tests {
		//quicksort(tt.args.nums, 0, len(tt.args.nums) - 1)
		quickSort(tt.args.nums)
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}



