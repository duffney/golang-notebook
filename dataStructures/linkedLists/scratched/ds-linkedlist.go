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

func (l *linkedList) Append(v int) {
	if l.length == 0 {
		l.head = &node{value: v}
		l.length++
		return
	}
	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = &node{value: v}
	l.length++
}

func (l *linkedList) append(n *node) {
	if l.length == 0 {
		l.head = n
		l.length++
		return
	}
	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = n
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

func (l *linkedList) prepend(n *node) {
	preHead := l.head
	l.head = n
	l.head.next = preHead
	l.length++
}

func (l *linkedList) Prepend(v int) {
	preHead := l.head
	l.head = &node{value: v}
	l.head.next = preHead
	l.length++
}

func (l *linkedList) deleteWithValue(v int) {
	if l.head == nil {
		return
	}
	if l.head.value == v {
		l.head = l.head.next
		l.length--
		return
	}
	currentNode := l.head
	for currentNode.next != nil {
		currentNode.next = currentNode.next.next
		return
	}
	currentNode = currentNode.next
	l.length--
}

// TODO update methods to accept value int

func main() {
	list := linkedList{}
	// list.prepend(2)
	// fmt.Println(list)
	node1 := &node{value: 1}
	node2 := &node{value: 2}
	node3 := &node{value: 3}
	node4 := &node{value: 4}
	list.append(node1)
	list.append(node2)
	list.append(node3)
	list.prepend(node4)
	list.print()
	fmt.Printf("\n")
	list.deleteWithValue(4)
	list.Prepend(5)
	list.print()

	// fmt.Printf("\n")

	// newList := linkedList{}
	// newList.Append(1)
	// newList.Append(2)
	// newList.Append(3)
	// newList.print()
}

/* tests to write
- add to null list
- append to non null list
*/
