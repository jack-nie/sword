package algo

import (
	"reflect"
	"testing"
)

func TestFindPathWithSum(t *testing.T) {
	root := constructBinaryTree([]int{10, 5, 4, 7, 12}, []int{4, 5, 7, 10, 12})

	for _, tt := range []struct {
		root   *TreeNode
		target int
		wanted [][]int
	}{
		{root, 22, [][]int{{10, 5, 7}, {10, 12}}},
	} {
		if got := findPathWithSum(tt.root, tt.target); !reflect.DeepEqual(tt.wanted, got) {
			t.Errorf("failed, expected %d, got %d", tt.wanted, got)
		}
	}
}
