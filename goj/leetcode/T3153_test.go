package leetcode

import (
	"testing"
)

func Test3153(t *testing.T) {
	testcase := [][]int{
		{1, 9, 1},
		{13, 23, 12},
		{10, 10, 10},
	}
	for _, v := range testcase {
		t.Log(sumDigitDifferences(v))
	}
}

func Test_sumDigitDifferences(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumDigitDifferences(tt.args.nums); got != tt.want {
				t.Errorf("sumDigitDifferences() = %v, want %v", got, tt.want)
			}
		})
	}
}
