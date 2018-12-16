package algo

func findPathWithSum(root *TreeNode, target int) []int {
	if root == nil {
		return nil
	}
	path := make([]int, 0)
	return findPath(root, &path, 0, target)
}

func findPath(root *TreeNode, path *[]int, sum, target int) []int {
	sum += root.Val
	*path = append(*path, root.Val)
	var isLeaf bool
	if root.Left == nil && root.Right == nil {
		isLeaf = true
	}
	if sum == target && isLeaf {
		return *path
	}
	if root.Left != nil {
		findPath(root.Left, path, sum, target)
	}
	if root.Right != nil {
		findPath(root.Right, path, sum, target)
	}
	*path = (*path)[: len(*path)-1 : len(*path)-1]
	return *path
}
