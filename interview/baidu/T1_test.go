package main

import (
	"reflect"
	"testing"
)

func Test_allSubset(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			args: args{
				input: []string{"1", "2", "3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allSubset(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}
