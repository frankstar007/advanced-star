/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/11/28
 * @contact frankstarye@tencent.com
 **/

package link_list

import (
	"fmt"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "merge two sorted list",
			args: args{
				list1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
				list2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},

			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); got != nil{

				w := tt.want
				g := got

				for w != nil && g != nil {
					if w.Val != g.Val {
						t.Errorf("want val:%d got val:%d", w.Val, g.Val)
					}
					fmt.Println(fmt.Sprintf("w:%d, g:%d", w.Val, g.Val))
					w = w.Next
					g = g.Next
				}
				if w != nil || g != nil {
					t.Errorf("want and got not equal")
				}
			}
		})
	}
}
