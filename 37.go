package algo

import (
	"io"
	"strconv"
)

func serialize(root *TreeNode, w *io.PipeWriter) {
	var stack []*TreeNode
	stack = append(stack, root)
	var node *TreeNode
	for len(stack) != 0 {
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			w.Write([]byte("$"))
		} else {
			w.Write([]byte(strconv.Itoa(node.Val)))
			stack = append(stack, node.Left)
			stack = append(stack, node.Right)
		}
	}
	w.Close()
}

func deserialize(r *io.PipeReader, root **TreeNode) {
	buf := make([]byte, 1)
	if _, err := r.Read(buf); err == nil {
		if val, err := strconv.Atoi(string(buf)); err == nil {
			*root = new(TreeNode)
			(*root).Val = val
			deserialize(r, &(*root).Left)
			deserialize(r, &(*root).Right)
		}
	}
}
