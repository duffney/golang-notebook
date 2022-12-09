package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// func reverseList(head *ListNode) *ListNode {
// 	var prev *ListNode
// 	for head != nil {
// 		head.Next, prev, head = prev, head, head.Next
// 	}
// 	return prev
// }

func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode = nil

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func NewLinkedList(val int, others ...int) *ListNode {
	ll := &ListNode{Val: val}
	current := ll
	for _, other := range others {
		current.Next = &ListNode{Val: other}
		current = current.Next
	}
	return ll
}

func (ll *ListNode) toArray() []int {
	vals := []int{}
	current := ll
	for current != nil {
		vals = append(vals, current.Val)
		current = current.Next
	}
	return vals
}

func main() {
	list := NewLinkedList(1, 2, 3)
	list = reverseList(list)
	fmt.Println(list.toArray())
}
