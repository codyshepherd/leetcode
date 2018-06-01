import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil && root.Left.Val >= root.Val {
		return false
	}

	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}

	return DFS(root.Left, math.MinInt64, root.Val) && DFS(root.Right, root.Val, math.MaxInt64)
}

func DFS(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	var left bool
	var right bool
	if root.Left == nil {
		left = true
	} else if root.Left.Val < root.Val && root.Left.Val < max && root.Left.Val > min {
		left = DFS(root.Left, min, root.Val)
	} else {
		return false
	}

	if left == false {
		return false
	}

	if root.Right == nil {
		right = true
	} else if root.Right.Val > root.Val && root.Right.Val < max && root.Right.Val > min {
		right = DFS(root.Right, root.Val, max)
	} else {
		return false
	}

	return left && right
}