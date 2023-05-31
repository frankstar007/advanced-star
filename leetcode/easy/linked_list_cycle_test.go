/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/27
 * @contact frankstarye@tencent.com
 **/

package easy

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}

	sub := &ListNode{
		Val: -4,
	}

	head := &ListNode{
		Val: 3,
	}

	two := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 0,
			Next: sub,
		},
	}
	sub.Next = two
	head.Next = two


	sec := &ListNode{
		Val: 1,
	}

	secp := &ListNode{
		Val: 2,
	}

	sec.Next = secp
	secp.Next = sec

	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test cycle",
			args: args{
				head: head,
			},
			want: true,
		},
		{
			name: "test sec",
			args: args{
				head: sec,
			},
			want: true,
		},
		{
			name: "test sec",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
					},
				},
			},
			want: false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
