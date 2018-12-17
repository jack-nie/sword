package algo

import "testing"

func TestMirrorRecursively(t *testing.T) {
	root1 := &TreeNode{nil, nil, 7}
	root1.Left = &TreeNode{nil, nil, 8}
	root1.Right = &TreeNode{nil, nil, 6}

	root2 := &TreeNode{nil, nil, 7}
	root2.Left = &TreeNode{nil, nil, 6}
	root2.Right = &TreeNode{nil, nil, 8}

	got := mirrorRecursively(root2)
	if got.Left.Val != 8 {
		t.Errorf("Failed, expected %d, got %d", root1.Left.Val, got.Left.Val)
	}
}
