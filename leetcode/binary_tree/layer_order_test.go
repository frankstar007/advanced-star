/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/15
 * @contact frankstarye@tencent.com
 **/

package binary_tree

import (
	"reflect"
	"testing"
)

func Test_layerOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "test for layer order 1",
			args: args{
				root : &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			want: [][]int{{1},{2,3},{4,5,6,7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := layerOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("layerOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
