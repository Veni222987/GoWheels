package leetcode

import "testing"

func Test(t *testing.T) {
	obj := Constructor()
	obj.Insert("apple")
	obj.Insert("app")
	param_2 := obj.Search("app")
	param_3 := obj.StartsWith("apple")
	t.Log(param_2, param_3)
}

func TestTrie_Insert(t *testing.T) {
	type fields struct {
		root *trieNode
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Trie{
				root: tt.fields.root,
			}
			this.Insert(tt.args.word)
		})
	}
}
