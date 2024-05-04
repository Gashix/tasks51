package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)

	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}

	return root
}

// Helper function to print the tree (in-order traversal)
func printTree(root *TreeNode) {
	if root != nil {
		printTree(root.Left)
		fmt.Print(root.Val, " ")
		printTree(root.Right)
	}
}

func main() {
	// Example 1
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}
	root1.Left.Right = &TreeNode{Val: 2}
	root1.Right.Left = &TreeNode{Val: 2}
	root1.Right.Right = &TreeNode{Val: 4}

	target1 := 2
	fmt.Println("Original tree (Example 1):")
	printTree(root1)
	fmt.Println()
	root1 = removeLeafNodes(root1, target1)
	fmt.Println("Modified tree (Example 1):")
	printTree(root1)
	fmt.Println()

	// Example 2
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 3}
	root2.Right = &TreeNode{Val: 3}
	root2.Left.Left = &TreeNode{Val: 3}
	root2.Left.Right = &TreeNode{Val: 2}

	target2 := 3
	fmt.Println("Original tree (Example 2):")
	printTree(root2)
	fmt.Println()
	root2 = removeLeafNodes(root2, target2)
	fmt.Println("Modified tree (Example 2):")
	printTree(root2)
	fmt.Println()

	// Example 3
	root3 := &TreeNode{Val: 1}
	root3.Left = &TreeNode{Val: 2}
	root3.Left.Right = &TreeNode{Val: 2}
	root3.Left.Right.Left = &TreeNode{Val: 2}

	target3 := 2
	fmt.Println("Original tree (Example 3):")
	printTree(root3)
	fmt.Println()
	root3 = removeLeafNodes(root3, target3)
	fmt.Println("Modified tree (Example 3):")
	printTree(root3)
	fmt.Println()
}
