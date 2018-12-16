package algo

import "testing"

func TestIsBalancedTree(t *testing.T) {
	root := &TreeNode{nil, nil, 5}
	root.Left = &TreeNode{nil, nil, 3}
	root.Left.Left = &TreeNode{nil, nil, 2}
	root.Left.Right = &TreeNode{nil, nil, 4}
	root.Right = &TreeNode{nil, nil, 7}
	root.Right.Left = &TreeNode{nil, nil, 6}
	root.Right.Right = &TreeNode{nil, nil, 8}

	root1 := &TreeNode{nil, nil, 5}
	root1.Left = &TreeNode{nil, nil, 3}
	root1.Left.Left = &TreeNode{nil, nil, 2}
	root1.Left.Right = &TreeNode{nil, nil, 4}
	root1.Right = &TreeNode{nil, nil, 7}
	root1.Right.Left = &TreeNode{nil, nil, 6}

	root2 := &TreeNode{nil, nil, 5}
	root2.Left = &TreeNode{nil, nil, 3}
	root2.Left.Left = &TreeNode{nil, nil, 2}
	root2.Left.Right = &TreeNode{nil, nil, 4}

	for _, v := range []struct {
		root   *TreeNode
		wanted bool
	}{
		{root, true},
		{root1, true},
		{root2, false},
		{nil, true},
	} {
		got := isBalancedTree(v.root)
		if got != v.wanted {
			t.Fatalf("Failed, expected %t, got %t", v.wanted, got)
		}
	}
}

func TestIsBalancedTreeTwo(t *testing.T) {
	root := &TreeNode{nil, nil, 5}
	root.Left = &TreeNode{nil, nil, 3}
	root.Left.Left = &TreeNode{nil, nil, 2}
	root.Left.Right = &TreeNode{nil, nil, 4}
	root.Right = &TreeNode{nil, nil, 7}
	root.Right.Left = &TreeNode{nil, nil, 6}
	root.Right.Right = &TreeNode{nil, nil, 8}

	root1 := &TreeNode{nil, nil, 5}
	root1.Left = &TreeNode{nil, nil, 3}
	root1.Left.Left = &TreeNode{nil, nil, 2}
	root1.Left.Right = &TreeNode{nil, nil, 4}
	root1.Right = &TreeNode{nil, nil, 7}
	root1.Right.Left = &TreeNode{nil, nil, 6}

	root2 := &TreeNode{nil, nil, 5}
	root2.Left = &TreeNode{nil, nil, 3}
	root2.Left.Left = &TreeNode{nil, nil, 2}
	root2.Left.Right = &TreeNode{nil, nil, 4}

	for _, v := range []struct {
		root   *TreeNode
		wanted bool
	}{
		{root, true},
		{root1, true},
		{root2, false},
		{nil, true},
	} {
		got := isBalancedTreeTwo(v.root)
		if got != v.wanted {
			t.Fatalf("Failed, expected %t, got %t", v.wanted, got)
		}
	}
}
