/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/15
 * @contact frankstarye@tencent.com
 **/

package arrays

import "testing"

func Test_minMaxGame(t *testing.T) {
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
			name: "test minMaxGame",
			args: args{
				nums: []int{1,3,5,2,4,8,2,2},
			},
			want: 1,
		},
		{
			name: "test 2",
			args: args{
				nums: []int{70,38,21,22},
			},
			want: 22,
		},
		{
			name: "test3",
			args: args{
				nums: []int{810,831,908,631,554,917,392,544},
			},
			want: 554,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMaxGame(tt.args.nums); got != tt.want {
				t.Errorf("minMaxGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
