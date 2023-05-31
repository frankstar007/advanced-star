/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/17
 * @contact frankstarye@tencent.com
 **/

package easy

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test for remove ele",
			args: args{
				nums: []int{3,2,2,3},
				val: 3,
			},
			want: 2,
		},
		{
			name: "test for 2",
			args: args{
				nums: []int{0,1,2,2,3,0,4,2},
				val: 2,
			},
			want: 5,
		},
		{
			name: "test 1",
			args: args{
				nums: []int{1},
				val: 1,
			},
			want: 0,
		},
		{
			name: "test 2",
			args: args{
				nums: []int{3,3},
				val: 3,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rmElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
