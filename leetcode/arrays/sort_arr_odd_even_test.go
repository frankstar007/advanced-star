/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/11/20
 * @contact frankstarye@tencent.com
 **/

package arrays

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParityII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test for sort arr odd and even",
			args: args{
				nums: []int{4,2,5,7},
			},
			want: []int{4,5,2,7},
		},
		{
			name: "test for sec test instance",
			args: args{
				nums: []int{4,1,2,1},
			},
			want: []int{4,1,2,1},
		},
		{
			name: "test for third test instance",
			args: args{
				nums: []int{3,1,4,2},
			},
			want: []int{2,1,4,3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParityII(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParityII() = %v, want %v", got, tt.want)
			}
		})
	}
}
