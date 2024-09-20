package baiduround2

import (
	"testing"
)

func Test_layerTraverse(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				root: &TreeNode{Val: 1,
					Left: &TreeNode{
						Val: 2,
					}, Right: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 4}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(layerTraverse(tt.args.root))
		})
	}
}
