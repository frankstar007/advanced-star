/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/19
 * @contact frankstarye@tencent.com
 **/

package easy

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		//{
		//	name: "test 1",
		//	args: args{
		//		nums1 : []int{1,2,3,0,0,0},
		//		nums2 : []int{2,5,6},
		//		m : 3,
		//		n : 3,
		//	},
		//	want: []int{1,2,2,3,5,6},
		//},
		//{
		//	name: "test 2",
		//	args: args{
		//		nums1 : []int{1},
		//		nums2 : []int{},
		//		m : 1,
		//		n : 0,
		//	},
		//	want: []int{1},
		//},
		//{
		//	name: "test 3",
		//	args: args{
		//		nums1 : []int{0},
		//		nums2 : []int{1},
		//		m : 0,
		//		n : 1,
		//	},
		//	want: []int{1},
		//
		//},
		//
		//{
		//	name: "test 3",
		//	args: args{
		//		nums1 : []int{4,5,6,0,0,0},
		//		nums2 : []int{1,2,3},
		//		m : 3,
		//		n : 3,
		//	},
		//	want: []int{1,2,3,4,5,6},
		//
		//},

		{
			name: "test error",
			args: args{
				nums1: []int{0,0,3,0,0,0,0,0,0},
				nums2: []int{-1,1,1,1,2,3},
				m : 3,
				n : 6,
			},
		},
	}
	for _, tt := range tests {
		nums := merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(nums)
		})
	}
}
