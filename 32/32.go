package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func printTreeFromTopToBottom(root *TreeNode) {
	if root == nil {
		return
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func main() {
	root := &TreeNode{nil, nil, 5}
	root.Left = &TreeNode{nil, nil, 3}
	root.Left.Left = &TreeNode{nil, nil, 2}
	root.Left.Right = &TreeNode{nil, nil, 4}
	root.Right = &TreeNode{nil, nil, 7}
	root.Right.Left = &TreeNode{nil, nil, 6}
	root.Right.Right = &TreeNode{nil, nil, 8}
	printTreeFromTopToBottom(root)
}
