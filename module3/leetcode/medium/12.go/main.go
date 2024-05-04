package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	dummy := &ListNode{Next: list1}
	prev := dummy
	for i := 0; i < a; i++ {
		prev = prev.Next
	}
	for i := 0; i < b-a+1; i++ {
		prev.Next = prev.Next.Next
	}
	tail := list2
	for tail.Next != nil {
		tail = tail.Next
	}
	prev.Next, tail.Next = list2, prev.Next
	return dummy.Next
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, " ")
		node = node.Next
	}
	fmt.Println()
}

func main() {
	list1 := &ListNode{Val: 10}
	list1.Next = &ListNode{Val: 1}
	list1.Next.Next = &ListNode{Val: 13}
	list1.Next.Next.Next = &ListNode{Val: 6}
	list1.Next.Next.Next.Next = &ListNode{Val: 9}
	list1.Next.Next.Next.Next.Next = &ListNode{Val: 5}

	list2 := &ListNode{Val: 1000000}
	list2.Next = &ListNode{Val: 1000001}
	list2.Next.Next = &ListNode{Val: 1000002}

	a, b := 3, 4

	result := mergeInBetween(list1, a, b, list2)
	printList(result) // Output: 10 1 13 1000000 1000001 1000002 5
}
