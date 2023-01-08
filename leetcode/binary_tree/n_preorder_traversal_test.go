/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/30
 * @contact frankstarye@tencent.com
 **/

package binary_tree

import (
	"reflect"
	"testing"
)

func Test_preorder(t *testing.T) {
	type args struct {
		root *Node
	}
	fc1 := &Node{
		Val: 5,
	}
	fc2 := &Node{
		Val: 6,
	}

	first := &Node{
		Val: 3,
		Children: []*Node{fc1,fc2},
	}

	sec := &Node{
		Val: 2,
	}

	third := &Node{
		Val: 4,
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test for n tree",
			args: args{
				root: &Node{
					Val: 1,
					Children: []*Node{
						first,
						sec,
						third,
					},
				},
			},
			want: []int{1,3,5,6,2,4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
