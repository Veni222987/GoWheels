package main

import "testing"

func Test_testCase1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testCase1()
		})
	}
}
