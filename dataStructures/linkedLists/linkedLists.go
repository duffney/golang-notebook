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

func (l *linkedList) append(v int) {
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

func (l linkedList) print() {
	fmt.Printf("Linked List: ")
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.value)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) prepend(v int) {
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

func main() {
	list := linkedList{}
	list.append(1)
	list.append(2)
	list.print()
	list.prepend(3)
	list.print()
	list.deleteWithValue(3)
	list.print()
}

/*
https://medium.com/backendarmy/linked-lists-in-go-f8a7b27a03b9
https://www.youtube.com/watch?v=1S0_-VxPLJo
https://go.dev/doc/effective_go#pointers_vs_values:~:text=This%20still%20requires%20the%20method%20to%20return%20the%20updated%20slice.%20We%20can%20eliminate%20that%20clumsiness%20by%20redefining%20the%20method%20to%20take%20a%20pointer%20to%20a%20ByteSlice%20as%20its%20receiver%2C%20so%20the%20method%20can%20overwrite%20the%20caller%27s%20slice.
https://www.youtube.com/watch?v=njTh_OwMljA
*/
