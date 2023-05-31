/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/5/30
 * @contact frankstarye@tencent.com
 **/

package medium

import (
	"fmt"
	"testing"
)

func Test_dfs(t *testing.T) {
	type args struct {
		root   *TreeNode
		result []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test for dfs",
			args: args{
				root: &TreeNode{
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
		{
			name: "test",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := dfs(tt.args.root, []int{3, 5})
			for _, r := range rs {
				fmt.Println(r)
			}
		})
	}
}

func Test_delNodes(t *testing.T) {
	type args struct {
		root    *TreeNode
		deletes []int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "test for del",
			args: args{
				root: &TreeNode{
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
				deletes: []int{3, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
