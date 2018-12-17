package algo

import "testing"

func TestHasSubTree(t *testing.T) {
	root := &TreeNode{nil, nil, 5}
	root.Left = &TreeNode{nil, nil, 3}
	root.Left.Left = &TreeNode{nil, nil, 2}
	root.Left.Right = &TreeNode{nil, nil, 4}
	root.Right = &TreeNode{nil, nil, 7}
	root.Right.Left = &TreeNode{nil, nil, 6}
	root.Right.Right = &TreeNode{nil, nil, 8}

	root2 := &TreeNode{nil, nil, 7}
	root2.Left = &TreeNode{nil, nil, 6}
	root2.Right = &TreeNode{nil, nil, 8}

	got := hasSubTree(root, root2)
	if got != true {
		t.Errorf("Failed, expected %t, got %t", false, got)
	}
}
