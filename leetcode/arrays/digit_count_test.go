/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/11
 * @contact frankstarye@tencent.com
 **/

package arrays

import "testing"

func Test_digitCount(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				num: "1210",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digitCount(tt.args.num); got != tt.want {
				t.Errorf("digitCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
