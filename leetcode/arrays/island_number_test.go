/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/16
 * @contact frankstarye@tencent.com
 **/

package arrays

import "testing"

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}

	g1 := [][]byte {
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},

	}


	g2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test for island number",
			args: args{
				grid: g1,
			},
			want: 1,
		},
		{
			name: "test case 2",
			args: args{
				grid: g2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
