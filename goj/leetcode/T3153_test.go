package leetcode

import (
	"testing"
)

func Test_sumDigitDifferences(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case1",
			args: args{nums: []int{1, 9, 1}},
			want: 0,
		},
		{
			name: "case2",
			args: args{nums: []int{13, 23, 12}},
			want: 6,
		},
		{
			name: "case3",
			args: args{nums: []int{10, 10, 10}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumDigitDifferences(tt.args.nums); got != tt.want {
				t.Errorf("sumDigitDifferences() = %v, want %v", got, tt.want)
			}
		})
	}
}
