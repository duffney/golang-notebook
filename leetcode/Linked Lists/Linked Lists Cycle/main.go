package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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

func main() {
	list := NewLinkedList(3, 2, 0, -4)
	list.Next.Next.Next.Next = list.Next // create cycle, link -4.Next to 2
	fmt.Println(hasCycle(list))
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow, fast = slow.Next, fast.Next.Next // fast.Next.Next could be nil check in loop
	}
	return false
}
