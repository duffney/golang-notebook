package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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
	list := NewLinkedList(1, 2, 3)
	fmt.Println(list.toArray())
}
