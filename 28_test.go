package algo

import "testing"

func TestIsSymmetrical(t *testing.T) {
	root := &TreeNode{nil, nil, 8}
	root.Left = &TreeNode{nil, nil, 6}
	root.Left.Left = &TreeNode{nil, nil, 5}
	root.Left.Right = &TreeNode{nil, nil, 7}
	root.Right = &TreeNode{nil, nil, 6}
	root.Right.Left = &TreeNode{nil, nil, 7}
	root.Right.Right = &TreeNode{nil, nil, 5}

	var root1 *TreeNode

	root2 := &TreeNode{nil, nil, 7}
	root2.Left = &TreeNode{nil, nil, 6}
	root2.Right = &TreeNode{nil, nil, 8}

	for _, v := range []struct {
		root   *TreeNode
		wanted bool
	}{
		{root, true},
		{root1, true},
		{root2, false},
	} {
		got := isSymmetrical(v.root)
		if got != v.wanted {
			t.Errorf("Failed, expected %t, got %t", v.wanted, got)
		}
	}
}
