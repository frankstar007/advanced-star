/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/15
 * @contact frankstarye@tencent.com
 **/

package sort

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				intervals: [][]int{{1,3},{2,6},{8,10},{15,18}},
			},
			want: [][]int{{1,6},{8,10},{15,18}},
		},
		{
			name: "test 2",
			args: args{
				intervals: [][]int{{1,4},{4,5}},
			},
			want: [][]int{{1,5}},
		},
		{
			name: "test 3",
			args: args{
				intervals: [][]int{{1,4},{0,4}},
			},
			want: [][]int{{0,4}},
		},
		{
			name: "test 4",
			args: args{
				intervals: [][]int{{1,4},{0,1}},
			},
			want: [][]int{{0,4}},
		},
		{
			name: "test 5",
			args: args{
				intervals: [][]int{{1,4},{2,3}},
			},
			want: [][]int{{1,4}},
		},

		{
			name: "test 6",
			args: args{
				intervals: [][]int{{1,4},{0,5}},
			},
			want: [][]int{{0,5}},
		},

		{
			name: "test 7",
			args: args{
				intervals: [][]int{{1,4},{0,0}},
			},
			want: [][]int{{0,0},{1,4}},
		},


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_qSort(t *testing.T) {
	type args struct {
		nums  []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "restudy qsort",
			args: args{
				nums: []int{40, 100, 1, 5, 25, 10},
				start: 0,
				end: 5,
			},
		},
	}
	for _, tt := range tests {
		sequence(tt.args.nums)
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}