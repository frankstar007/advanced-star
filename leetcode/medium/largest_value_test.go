/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/5/23
 * @contact frankstarye@tencent.com
 **/

package medium

import "testing"

func Test_largestValFromLabels(t *testing.T) {
	type args struct {
		values    []int
		labels    []int
		numWanted int
		useLimit  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test for largest",
			args: args{
				values:    []int{9, 8, 8, 7, 6},
				labels:    []int{0, 0, 0, 1, 1},
				numWanted: 3,
				useLimit:  1,
			},
			want: 16,
		},
		{
			name: "test for sec",
			args: args{
				values:    []int{5, 4, 3, 2, 1},
				labels:    []int{1, 3, 3, 3, 2},
				numWanted: 3,
				useLimit:  2,
			},
			want: 12,
		},
		{
			name: "test for third",
			args: args{
				values:    []int{5, 4, 3, 2, 1},
				labels:    []int{1, 1, 2, 2, 3},
				numWanted: 3,
				useLimit:  1,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestValFromLabels(tt.args.values, tt.args.labels, tt.args.numWanted, tt.args.useLimit); got != tt.want {
				t.Errorf("largestValFromLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
