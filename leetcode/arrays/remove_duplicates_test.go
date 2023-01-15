/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/10
 * @contact frankstarye@tencent.com
 **/

package arrays

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				nums: []int{1,1,2},
			},
			want: 2,
		},
		{
			name: "test 2",
			args: args{
				nums: []int{0,0,1,1,1,2,2,3,3,4},
			},
			want: 5,
		},
		{
			name: "test 3",
			args: args{
				nums: []int{1,1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := removeDuplicates(tt.args.nums)
			if got1 != tt.want {
				t.Errorf("removeDuplicateszzzzz() = %v, want %v", got1, tt.want)
			}

		})
	}
}
