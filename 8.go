package algo

func getNext(root *TreeNodeWithParent) *TreeNodeWithParent {
	if root == nil {
		return nil
	}
	var next *TreeNodeWithParent
	if root.Right != nil {
		right := root.Right
		for right.Left != nil {
			right = right.Left
			next = right
		}
	} else if root.Parent != nil {
		current := root
		parent := current.Parent
		for parent != nil && current == parent.Right {
			current = parent
			parent = parent.Parent
		}
		next = parent
	}
	return next
}
