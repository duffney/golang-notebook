package main

import "fmt"

/*
type, just ust int this time...
new func
methods
- insert
- delete
- max
- find
- print/traverse
*/

type SortedArray struct {
	Elements []int
	Length   int
	Capacity int
}

func NewSortedArray(capacity int) SortedArray {
	return SortedArray{
		Elements: make([]int, capacity),
		Length:   0,
		Capacity: capacity,
	}
}

// insertion sort
func (a *SortedArray) Insert(value int) bool {
	if a.Length >= a.Capacity {
		return false // Return false if the array is full
	}

	// Start from the last occupied index
	i := a.Length - 1

	// Shift elements to the right until the correct position for 'e' is found
	for i >= 0 && a.Elements[i] > value {
		a.Elements[i+1] = a.Elements[i] // Move the current element one position to the right
		i--                             // Move left to check the next element
	}

	// Insert 'e' at the found position (i+1)
	a.Elements[i+1] = value
	a.Length++

	return true // Return true to indicate successful insertion
}

// To remove an element from a sorted array,
// first we find its index,
// then we shift all the elements to its right left one position, overwriting the deleted element.
func (a *SortedArray) Delete(target int) bool {
	var index int
	for i, e := range a.Elements {
		if e == target {
			index = i
		}
	}

	// Loop through starting at the index
	for i := index; i < a.Length-1; i++ {
		//Shift elements left one position
		a.Elements[i] = a.Elements[i+1]
	}
	// Blank out last element in the array
	a.Elements[a.Length-1] = 0
	a.Length--

	return true
}

func (a *SortedArray) LinearSearch(target int) bool {
	for i := 0; i < a.Length; i++ {
		if a.Elements[i] == target {
			fmt.Printf("found, %d at index %d\n", a.Elements[i], i)
			return true
		}
	}
	return false
}

func (a *SortedArray) Find(target int) bool {
	left := 0
	right := a.Length - 1
	return a.BinarySearch(target, left, right)
}

// BinarySearch performs a recursive binary search
func (a *SortedArray) BinarySearch(target, left, right int) bool {
	// Base case 1: Search space exhausted (element not found)
	if left > right {
		return false
	}

	// Correct midpoint calculation
	mid := left + (right-left)/2

	// Base case 2: Target found
	if a.Elements[mid] == target {
		fmt.Printf("found %d at index %d\n", a.Elements[mid], mid)
		return true
	}

	// Recursive cases
	if a.Elements[mid] > target {
		// Search left half and return the result
		return a.BinarySearch(target, left, mid-1)
	}
	// Note: No else needed; if > fails, it must be <
	// Search right half and return the result
	return a.BinarySearch(target, mid+1, right)
}

func (a *SortedArray) BinarySearchNonRecursive(target int) bool {
	left := 0
	right := a.Length - 1

	for left <= right {
		mid := left + (right-left)/2

		if a.Elements[mid] == target {
			fmt.Printf("found %d at index %d\n", a.Elements[mid], mid)
			return true
		}

		if a.Elements[mid] > target {
			right = mid + 1
		} else {
			left = mid + 1
		}
	}
	return false
}
