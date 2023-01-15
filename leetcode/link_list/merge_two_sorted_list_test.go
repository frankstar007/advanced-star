/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/9
 * @contact frankstarye@tencent.com
 **/

package link_list

import (
	"fmt"
	"testing"
)

func Test_mergeTwoSortedLists(t *testing.T) {
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
			name : "test 1",
			args: args{
				list1 :&ListNode{
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
		{
			name: "test 2",
			args: args{
				list1: &ListNode{
					Val: 2,
				},
				list2: &ListNode{
					Val: 1,
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//got := mergeTwoSortedListsWithNew(tt.args.list1, tt.args.list2)
			got1 := mergeTwoSortedLists(tt.args.list1, tt.args.list2)
			//fmt.Println(got.ToString())
			fmt.Println(got1.ToString())
			if !tt.want.Equal(got1){
				t.Errorf("mergeTwoSortedLists() = %v, want %v", got1, tt.want)
			}

		})
	}
}
