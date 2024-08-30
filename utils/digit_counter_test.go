package utils

import "testing"

func TestDigitCounter(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{num: 123}, want: 3},
		{args: args{num: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DigitCounter(tt.args.num); got != tt.want {
				t.Errorf("DigitCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}
