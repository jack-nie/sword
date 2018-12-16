package algo

import "testing"

func TestFindKthNode(t *testing.T) {
	root := &TreeNode{nil, nil, 5}
	root.Left = &TreeNode{nil, nil, 3}
	root.Left.Left = &TreeNode{nil, nil, 2}
	root.Left.Right = &TreeNode{nil, nil, 4}
	root.Right = &TreeNode{nil, nil, 7}
	root.Right.Left = &TreeNode{nil, nil, 6}
	root.Right.Right = &TreeNode{nil, nil, 8}
	var k int = 3
	node := findKthNode(root, &k)
	if node.Val != 4 {
		t.Errorf("Failed, expected %d, got %d", 4, node.Val)
	}

}
