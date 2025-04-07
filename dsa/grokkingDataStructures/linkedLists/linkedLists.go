package main

import (
	"fmt"
	"reflect"
)

type Song struct {
	Name   string
	Artist string
	Album  string
}

type node struct {
	value Song
	next  *node
}

type LinkedList struct {
	head *node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil, size: 0}
}

func (l *LinkedList) Append(value Song) {
	newNode := &node{
		value: value,
		next:  l.head,
	}
	l.head = newNode
	l.size++
}

func (l *LinkedList) Display() {
	current := l.head

	for current != nil {
		fmt.Printf("Name: %s, Artist: %s, Album: %s\n", current.value.Name, current.value.Artist, current.value.Album)
		current = current.next
	}
}
func (l *LinkedList) Search(value, field string) Song {
	current := l.head

	for current != nil {
		v := reflect.ValueOf(current.value)
		f := v.FieldByName(field)

		// Check if the field exists and is a string
		if f.IsValid() && f.Kind() == reflect.String && f.String() == value {
			return current.value
		}
		current = current.next
	}
	return Song{}
}

// TODO: Delete, handle edge cases
// If you delete the last node, then you have to make the previous node the new tail (in our implementation, its link is set to None).
// If you delete the first node, the head of the list, then there is no predecessor! In this case, we just have to update the list’s head pointer.
func (l *LinkedList) DeleteSong(value string) bool {
	current := l.head
	var prev *node

	for current != nil {
		if current.value.Name == value {
			if prev == nil {
				l.head = current.next
			} else {
				// set previous.Next to current.next
				prev.next = current.next
			}
			l.size--
			return true
		}
		prev = current
		current = current.next
	}
	return false
}

func main() {
	list := NewLinkedList()

	list.Append(Song{Name: "My Own Prison", Artist: "Creed", Album: "My Own Prison"})
	list.Append(Song{Name: "Higher", Artist: "Creed", Album: "Human Clay"})
	list.Append(Song{Name: "With Arms Wide Open", Artist: "Creed", Album: "Human Clay"})

	// list.Display()

	// s := list.Search("Higher", "Name")
	// fmt.Println(s)
	list.DeleteSong("With Arms Wide Open")
	list.Display()
}
