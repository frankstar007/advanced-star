/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/19
 * @contact frankstarye@tencent.com
 **/

package easy

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				digits: []int{8,9,9,9},
			},
			want: []int{9, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
