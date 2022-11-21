/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/11/19
 * @contact frankstarye@tencent.com
 **/

package dynamic_programming

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name : "test for sell stock for max profit",
			args: args{
				prices: []int{7,1,5,3,6,4},
			},
			want: 5,
		},
		{
			name: "test for sell stock for max profit sec",
			args: args{
				prices: []int{7,6,4,3,1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dpMaxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name : "dp: test for sell stock for max profit",
			args: args{
				prices: []int{7,1,5,3,6,4},
			},
			want: 5,
		},
		{
			name: "dp: test for sell stock for max profit sec",
			args: args{
				prices: []int{7,6,4,3,1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dpMaxProfit(tt.args.prices); got != tt.want {
				t.Errorf("dpMaxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}