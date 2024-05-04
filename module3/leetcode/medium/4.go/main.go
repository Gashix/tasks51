package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	traverse(root, &sum)
	return root
}

func traverse(node *TreeNode, sum *int) {
	if node == nil {
		return
	}

	traverse(node.Right, sum)
	*sum += node.Val
	node.Val = *sum
	traverse(node.Left, sum)
}

func printInOrder(root *TreeNode) {
	if root == nil {
		return
	}
	printInOrder(root.Left)
	fmt.Printf("%d ", root.Val)
	printInOrder(root.Right)
}

func main() {
	root1 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 7,
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}
	fmt.Println("Original BST:")
	printInOrder(root1)
	fmt.Println()

	convertedRoot1 := bstToGst(root1)
	fmt.Println("Converted Greater Tree:")
	printInOrder(convertedRoot1)
	fmt.Println()

	root2 := &TreeNode{
		Val: 0,
		Right: &TreeNode{
			Val: 1,
		},
	}
	fmt.Println("Original BST:")
	printInOrder(root2)
	fmt.Println()

	convertedRoot2 := bstToGst(root2)
	fmt.Println("Converted Greater Tree:")
	printInOrder(convertedRoot2)
	fmt.Println()
}
