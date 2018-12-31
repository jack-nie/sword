package algo

func findPathWithSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	path := make([]int, 0)
	result := make([][]int, 0)
	findPath(root, &path, &result, 0, target)
	return result
}

func findPath(root *TreeNode, path *[]int, result *[][]int, sum, target int) {
	sum += root.Val
	*path = append(*path, root.Val)
	var isLeaf bool
	if root.Left == nil && root.Right == nil {
		isLeaf = true
	}
	if sum == target && isLeaf {
		*result = append(*result, *path)
		*path = (*path)[: len(*path)-1 : len(*path)-1]
		return
	}
	if root.Left != nil {
		findPath(root.Left, path, result, sum, target)
	}
	if root.Right != nil {
		findPath(root.Right, path, result, sum, target)
	}
	*path = (*path)[: len(*path)-1 : len(*path)-1]
}
