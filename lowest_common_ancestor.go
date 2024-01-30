package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return check(root, p, q)
}

func check(node, p, q *TreeNode) *TreeNode {
	if node.Left == nil {
		return nil
	}

	if node.Right == nil {
		return nil
	}

	// check if its self descendant
	if node.Val == p.Val || node.Val == q.Val {
		if (node.Left.Val == p.Val || node.Left.Val == q.Val) && (node.Right.Val == p.Val || node.Right.Val == q.Val) {
			return node
		}
	}

	// check children
	if (node.Left.Val == p.Val || node.Left.Val == q.Val) && (node.Right.Val == p.Val || node.Right.Val == q.Val) {
		return node
	}

	return nil
}
