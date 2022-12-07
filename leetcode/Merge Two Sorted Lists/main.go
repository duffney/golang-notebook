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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	h := &ListNode{} // dummy head with nul to avoid having to create a head value
	l := h
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			l.Next = list1
			list1 = list1.Next
		} else {
			l.Next = list2
			list2 = list2.Next
		}
		l = l.Next // create next node in Linked List
	}

	if list1 == nil {
		l.Next = list2 // append the remaining list2 nodes
	}

	if list2 == nil {
		l.Next = list1 // append the remaining list1 nodes
	}

	return h.Next
}

func (ll *ListNode) ToArray() []int {
	vals := []int{}
	current := ll
	for current != nil {
		vals = append(vals, current.Val)
		current = current.Next
	}
	return vals
}

func main() {
	list1 := NewLinkedList(1, 2, 4)
	list2 := NewLinkedList(1, 3, 4)
	output := mergeTwoLists(list1, list2)
	fmt.Println(output.toArray())
}
