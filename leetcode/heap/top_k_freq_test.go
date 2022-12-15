/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/4
 * @contact frankstarye@tencent.com
 **/

package heap

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "top k frequent",
			args: args{
				nums: []int{1,1,1,2,2,3},
				k : 2,
			},
			want: []int{1,2},
		},
		{
			name: "top k frequent",
			args: args{
				nums: []int{1},
				k : 1,
			},
			want: []int{1},
		},
		{
			name: "test case 3",
			args: args{
				nums: []int{4,1,-1,2,-1,2,3},
				k : 2,
			},
			want: []int{-1,2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
