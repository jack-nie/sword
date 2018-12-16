package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func printTreeByZigzag(root *TreeNode) {
	if root == nil {
		return
	}
	stack1 := make([]*TreeNode, 0)
	stack2 := make([]*TreeNode, 0)
	level := 1
	stack1 = append(stack1, root)

	for len(stack1) != 0 || len(stack2) != 0 {
		if level&0x1 != 0 {
			for len(stack1) != 0 {
				node := stack1[len(stack1)-1]
				stack1 = stack1[:len(stack1)-1]
				if node != nil {
					fmt.Print(node.Val)
					if node.Left != nil {
						stack2 = append(stack2, node.Left)
					}
					if node.Right != nil {
						stack2 = append(stack2, node.Right)
					}
				}

			}
			level++
			fmt.Print("\n")
		} else {
			for len(stack2) != 0 {
				node := stack2[len(stack2)-1]
				stack2 = stack2[:len(stack2)-1]
				if node != nil {
					fmt.Print(node.Val)
					if node.Right != nil {
						stack1 = append(stack1, node.Right)
					}
					if node.Left != nil {
						stack1 = append(stack1, node.Left)
					}
				}

			}
			level++
			fmt.Print("\n")
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
	printTreeByZigzag(root)
}
