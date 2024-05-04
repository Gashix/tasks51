package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var dfs func(node *TreeNode) (int, int) // sum, count
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		leftSum, leftCount := dfs(node.Left)
		rightSum, rightCount := dfs(node.Right)
		nodeSum := node.Val + leftSum + rightSum
		nodeCount := 1 + leftCount + rightCount
		return nodeSum, nodeCount
	}

	var countEqual func(node *TreeNode, avg int) int
	countEqual = func(node *TreeNode, avg int) int {
		if node == nil {
			return 0
		}
		res := 0
		if node.Val == avg {
			res++
		}
		leftRes := countEqual(node.Left, avg)
		rightRes := countEqual(node.Right, avg)
		return res + leftRes + rightRes
	}

	var ans int
	var dfsAndCheck func(node *TreeNode)
	dfsAndCheck = func(node *TreeNode) {
		if node == nil {
			return
		}
		leftSum, leftCount := dfs(node.Left)
		rightSum, rightCount := dfs(node.Right)
		nodeSum := node.Val + leftSum + rightSum
		nodeCount := 1 + leftCount + rightCount
		avg := nodeSum / nodeCount
		ans += countEqual(node, avg)
		dfsAndCheck(node.Left)
		dfsAndCheck(node.Right)
	}

	dfsAndCheck(root)
	return ans
}

func main() {
	// Example usage:
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 8}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 6}

	fmt.Println(averageOfSubtree(root)) // Output: 5
}
