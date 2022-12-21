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
	list := NewLinkedList(1, 2, 3, 4, 5)
	middle := middleNode(list)
	fmt.Println(middle.toArray())
}

func middleNode(head *ListNode) *ListNode {
	array := []*ListNode{}
	current := head
	for current != nil {
		array = append(array, current)
		current = current.Next
	}

	return array[len(array)/2]
}
