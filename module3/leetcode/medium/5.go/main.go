package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	// Find the length of the linked list
	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}

	// Split the linked list into two halves
	curr = head
	for i := 0; i < length/2-1; i++ {
		curr = curr.Next
	}
	secondHalfStart := curr.Next
	curr.Next = nil

	// Reverse the second half of the linked list
	secondHalfStart = reverseLinkedList(secondHalfStart)

	// Iterate through the first half and the reversed second half to find the maximum twin sum
	maxSum := 0
	for head != nil {
		twinSum := head.Val + secondHalfStart.Val
		if twinSum > maxSum {
			maxSum = twinSum
		}
		head = head.Next
		secondHalfStart = secondHalfStart.Next
	}

	return maxSum
}

func reverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func createLinkedList(arr []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, val := range arr {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}
	return dummy.Next
}

func main() {
	head1 := createLinkedList([]int{5, 4, 2, 1})
	fmt.Println(pairSum(head1)) // Output: 6

	head2 := createLinkedList([]int{4, 2, 2, 3})
	fmt.Println(pairSum(head2)) // Output: 7

	head3 := createLinkedList([]int{1, 100000})
	fmt.Println(pairSum(head3)) // Output: 100001
}
