package main

import "fmt"

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) add(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) print() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.value)
		toPrint = toPrint.next
		l.length--
	}
}

func (l *linkedList) delete(value int) {
	if l.length == 0 { // Was an empty list passed?
		return
	}

	if l.head.value == value { // Is the head being deleted?
		l.head = l.head.next
		l.length--
		return
	}
	node := l.head
	for node.next.value != value {
		if node.next.next == nil {
			return
		}
		node = node.next
		//fmt.Printf("%d ", l.head.value)
	}
	node.next = node.next.next
	l.length--
}

func main() {
	l1 := linkedList{}
	node1 := &node{value: 1}
	node2 := &node{value: 2}
	node3 := &node{value: 3}
	l1.add(node1)
	l1.add(node2)
	l1.add(node3)
	l1.delete(100)
	l1.print()

	// l2 := linkedList{}
	// l2node := &node{value: 1}
	// l2.add(l2node)

	// x := 7
	// y := &x
	// fmt.Println(*y)
}
