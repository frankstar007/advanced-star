/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/30
 * @contact frankstarye@tencent.com
 **/

package binary_tree

import "testing"

func Test_inorderRecursion(t *testing.T) {
	type args struct {
		node *TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "inorder traversal test",
			args: args{
				node: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
					},
				},
			},
		},
		{
			name : " inorder traversal iteration test",
			args: args{
				node : &TreeNode{
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inorderTraversal(tt.args.node)
			inorderIteration(tt.args.node)
		})
	}
}
