/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/11/19
 * @contact frankstarye@tencent.com
 **/

package link_list

import (
	"reflect"
	"testing"
)

func Test_iteratorReverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "iterator reverse link list",

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := iteratorReverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("iteratorReverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "cycle link list",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
