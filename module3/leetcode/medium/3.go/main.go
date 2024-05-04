package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeZeroNodes(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy
	sum := 0

	for head != nil {
		if head.Val != 0 {
			sum += head.Val
		} else {
			if sum != 0 {
				prev.Next = &ListNode{Val: sum, Next: nil}
				prev = prev.Next
			}
			sum = 0
		}
		head = head.Next
	}

	if sum != 0 {
		prev.Next = &ListNode{Val: sum, Next: nil}
	}

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	head := &ListNode{Val: 0, Next: &ListNode{Val: 3, Next: &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0}}}}}}}}
	result := mergeZeroNodes(head)
	printList(result) // Output: 4 11
}
