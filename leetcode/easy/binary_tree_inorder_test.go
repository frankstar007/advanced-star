/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/23
 * @contact frankstarye@tencent.com
 **/

package easy

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test recursive inorder binary tree",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
					},
				},
			},
			want: []int{1,3,2},
		},
		{
			name: "test case",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 7,
							},
							Right: &TreeNode{
								Val: 8,
							},
						},
						Right: &TreeNode{
							Val: 6,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: []int{7, 4, 8, 2, 6, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := iterateInorder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
			if got := colorInorder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
