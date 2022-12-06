package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := NewNode(1)
	l1.Next = NewNode(2)
	l1.Next.Next = NewNode(4)

	l2 := NewNode(1)
	l2.Next = NewNode(3)
	l2.Next.Next = NewNode(4)
	//TraverseLinkedList(l1)
	fmt.Println(mergeTwoLists(l1, l2))
}

func NewNode(value int) *ListNode {
	var n ListNode
	n.Val = value
	n.Next = nil
	return &n
}

func TraverseLinkedList(head *ListNode) {
	fmt.Printf("Linked List: ")
	temp := head
	for temp != nil {
		fmt.Printf("%d ", temp.Val)
		temp = temp.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	return list2
}
