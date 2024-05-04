package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2

	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}

func inorderTraversal(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left, values)
	*values = append(*values, root.Val)
	inorderTraversal(root.Right, values)
}

func balanceBST(root *TreeNode) *TreeNode {
	var values []int
	inorderTraversal(root, &values)
	return sortedArrayToBST(values)
}

func printTreeInOrder(root *TreeNode) {
	if root == nil {
		return
	}
	printTreeInOrder(root.Left)
	fmt.Printf("%d ", root.Val)
	printTreeInOrder(root.Right)
}

func main() {
	root1 := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}},
	}
	fmt.Println("Original tree:")
	printTreeInOrder(root1)
	fmt.Println()

	balancedTree1 := balanceBST(root1)
	fmt.Println("Balanced tree:")
	printTreeInOrder(balancedTree1)
	fmt.Println()

	root2 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	fmt.Println("Original tree:")
	printTreeInOrder(root2)
	fmt.Println()

	balancedTree2 := balanceBST(root2)
	fmt.Println("Balanced tree:")
	printTreeInOrder(balancedTree2)
	fmt.Println()
}
