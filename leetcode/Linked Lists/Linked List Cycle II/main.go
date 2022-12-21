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
	list := NewLinkedList(3, 2, 0, -4)
	list.Next.Next.Next.Next = list.Next // create cycle at list index 1
	fmt.Println(detectCycle(list))
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast, isCycle := head, head, false

	for fast != nil && fast.Next != nil { //detect cycle logic
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			isCycle = true
			break
		}
	}

	if !isCycle { // if there isn't a cycle return nil
		return nil
	}

	slow = head        //reset slow to head
	for slow != fast { // find cycle node in linkedlist (which equals fast)
		slow, fast = slow.Next, fast.Next
	}

	return slow
}

/*
https://www.youtube.com/watch?v=MFOAbpfrJ8g
https://leetcode.com/problems/linked-list-cycle-ii/solutions/1183963/floyd-s-algorithm-go-golang/
*/
