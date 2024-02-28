package RootEqualSum

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	sum := root.Left.Val + root.Right.Val
	if sum != root.Val {
		return false
	} else {
		return true
	}
}
