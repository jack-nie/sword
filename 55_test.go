package algo

import "testing"

func TestTreeDepth(t *testing.T) {
	root := new(TreeNode)
	root.Left = new(TreeNode)
	root.Right = new(TreeNode)
	root.Left.Left = new(TreeNode)
	root.Left.Left = new(TreeNode)
	depth := treeDepth(root)
	if depth != 3 {
		t.Errorf("Failed, expect %d, got %d", 3, depth)
	}
}
