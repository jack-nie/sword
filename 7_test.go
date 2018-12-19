package algo

import (
	"reflect"
	"testing"
)

func TestConstructBinaryTree(t *testing.T) {
	root := &TreeNode{nil, nil, 1}
	root.Left = &TreeNode{nil, nil, 2}
	root.Left.Left = &TreeNode{nil, nil, 4}
	root.Left.Left.Right = &TreeNode{nil, nil, 7}
	root.Right = &TreeNode{nil, nil, 3}
	root.Right.Left = &TreeNode{nil, nil, 5}
	root.Right.Right = &TreeNode{nil, nil, 6}
	root.Right.Right.Left = &TreeNode{nil, nil, 8}

	for _, v := range []struct {
		preOrder []int
		inOrder  []int
		wanted   *TreeNode
	}{
		{[]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6}, root},
	} {
		var stack1, stack2 []int
		got := inOrder(constructBinaryTree(v.preOrder, v.inOrder), &stack1)
		wanted := inOrder(v.wanted, &stack2)
		if !reflect.DeepEqual(got, wanted) {
			t.Errorf("Failed, expected %d, got %d", wanted, got)
		}

	}
}

func inOrder(root *TreeNode, stack *[]int) []int {
	if root != nil {
		inOrder(root.Left, stack)
		*stack = append(*stack, root.Val)
		inOrder(root.Right, stack)
	}
	return *stack
}
