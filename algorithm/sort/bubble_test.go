/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/10
 * @contact frankstarye@tencent.com
 **/

package sort

import "testing"

func Test_bubbleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test for bubble",
			args: args{
				nums: []int{20,10,40,30,10,60,50},
			},
		},
	}
	for _, tt := range tests {
		bubbleSort(tt.args.nums)
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
