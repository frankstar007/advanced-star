/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/5/29
 * @contact frankstarye@tencent.com
 **/

package graph

import (
	"fmt"
	"testing"
)

func TestNewGraph(t *testing.T) {
	type args struct {
		grid [][]int
	}

	g := [][]int{
		{1, 0},
		{0, 1},
	}

	g1 := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}

	tests := []struct {
		name string
		args args
		want Graph
	}{
		// TODO: Add test cases.
		{
			name: "build graph",
			args: args{
				grid: g,
			},
			want: Graph{},
		},
		{
			name: "build graph",
			args: args{
				grid: g1,
			},
			want: Graph{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGraph(tt.args.grid)
			result := got.Traverse()
			fmt.Println(result)
			dis := got.ShortPath()
			fmt.Println(dis)

		})
	}
}
