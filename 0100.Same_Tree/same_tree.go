package _100_Same_Tree

// TreeNode is a definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Equal nullity denotes that this branch is the same (local equality)
	// This is a base case, but also handles being given two empty trees
	if p == nil && q == nil {
		return true
	}

	// Unequal nullity denotes that the trees aren't the same
	// Note that we've already ruled out equal nullity above
	if p == nil || q == nil {
		return false
	}

	// Both nodes have values; descend if those values are equal
	// "&&" here allows for any difference to overrule a local equality
	if p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}

	// If we're here, both nodes have values, and they're unequal, so the trees aren't the same
	return false
}
