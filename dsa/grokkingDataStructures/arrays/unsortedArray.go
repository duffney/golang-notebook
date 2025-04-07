package main

import "fmt"

type UnsortedArray[T ~int | ~string] struct {
	Elements []T
	Length   int
	Capacity int
}

func NewUnsortedArray[T ~int | ~string](capactiy int) UnsortedArray[T] {
	return UnsortedArray[T]{
		Elements: make([]T, capactiy),
		Length:   0,
		Capacity: capactiy,
	}
}

func (a *UnsortedArray[T]) Insert(e T) bool {
	if a.Length >= a.Capacity {
		return false
	}
	a.Elements[a.Length] = e
	a.Length++
	return true
}

// Move selected element to the end, then delete
func (a *UnsortedArray[T]) DeleteByIndex(index int) bool {
	if a.Length == 0 {
		return false
	}

	if index < 0 || index >= a.Capacity {
		return false
	}
	a.Elements[index] = a.Elements[a.Length-1]
	a.Length--
	return true
}

func (a *UnsortedArray[T]) Find(target T) bool {
	for i := range a.Length {
		if a.Elements[i] == target {
			fmt.Printf("target: %v found at index: %d\n", target, i)
			return true
		}
	}
	return false
}

func (a *UnsortedArray[T]) Print() {
	fmt.Println(a.Elements[:a.Length])
}

// can you make a generic method work for only 1 of the types? ex int not string
func (a *UnsortedArray[T]) Max() T {
	var max T
	for i := range a.Length {
		fmt.Printf("current: %v i: %v\n", max, a.Elements[i])
		if max < a.Elements[i] {
			max = a.Elements[i]
		}
	}
	return max
}
