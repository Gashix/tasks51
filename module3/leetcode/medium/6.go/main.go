package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxIndex := 0
	for i, num := range nums {
		if num > nums[maxIndex] {
			maxIndex = i
		}
	}

	root := &TreeNode{Val: nums[maxIndex]}

	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])

	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	nums1 := []int{3, 2, 1, 6, 0, 5}
	tree1 := constructMaximumBinaryTree(nums1)
	printTree(tree1)
	// Output: 6 3 2 1 5 0

	fmt.Println()

	nums2 := []int{3, 2, 1}
	tree2 := constructMaximumBinaryTree(nums2)
	printTree(tree2)
	// Output: 3 2 1
}
