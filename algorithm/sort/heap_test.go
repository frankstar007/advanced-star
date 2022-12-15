/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/4
 * @contact frankstarye@tencent.com
 **/

package sort

import (
	"fmt"
	"testing"
)

func Test_heapify(t *testing.T) {
	type args struct {
		tree  []int
		n     int
		index int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test for heapify with the defined node",
			args: args{
				tree: []int{4, 10, 3, 5, 1, 2},
				n : 6,
				index: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("before heapify:%v\n", tt.args.tree)
			heapify(tt.args.tree, tt.args.n, tt.args.index)
			fmt.Printf("after heapify:%v\n", tt.args.tree)
		})

	}
}

func Test_build_heap(t *testing.T) {
	type args struct {
		tree []int
		n    int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "build heap tree",
			args: args{
				tree: []int{1,2,3,4,5,6,7},
				n: 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("before build :%v\n", tt.args.tree)
			buildHeap(tt.args.tree, tt.args.n)
			fmt.Printf("after build :%v\n", tt.args.tree)
		})
	}
}

func Test_heapSort(t *testing.T) {
	type args struct {
		tree []int
		n    int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test for heap sort",
			args: args{
				tree: []int{3,2,3,1,2,4,5,5,6},
				n : 9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heapSort(tt.args.tree, tt.args.n)
		})
	}
}