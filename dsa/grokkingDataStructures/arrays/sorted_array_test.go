package main

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	// Test case 1: Insert into an empty sorted array
	t.Run("Insert into empty array", func(t *testing.T) {
		arr := NewSortedArray(3)
		success := arr.Insert(10)

		if !success {
			t.Errorf("Insert failed when the array is empty")
		}

		if arr.Length != 1 {
			t.Errorf("Expected Length to be 1, got %d", arr.Length)
		}

		if arr.Elements[0] != 10 {
			t.Errorf("Expected first element to be 10, got %d", arr.Elements[0])
		}
	})
	t.Run("Insert into empty unsorted array", func(t *testing.T) {
		ua := NewUnsortedArray[int](3)
		success := ua.Insert(2)

		if !success {
			t.Errorf("Insert faield when the array is empty")
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("Insert into empty array", func(t *testing.T) {
		sa := NewSortedArray(4)
		ints := []int{1, 2, 3, 4}
		for i := range ints {
			sa.Insert(i)
		}

		sucess := sa.Delete(2)

		if !sucess {
			t.Errorf("Delete failed")
		}
	})
}

func TestInsertAlgorithms(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		insert   int
		expected []int
	}{
		{
			name:     "Insert into empty array",
			initial:  []int{},
			insert:   1,
			expected: []int{1},
		},
		{
			name:     "Insert 2 between 1,3",
			initial:  []int{1, 3},
			insert:   2,
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		// Test the original algorithm
		t.Run("Original: "+tt.name, func(t *testing.T) {
			// Create and initialize array
			a := NewSortedArray(10)
			a.Length = len(tt.initial)
			for val := range tt.initial {
				a.Insert(val)
			}

			// Insert the value
			a.Insert(tt.insert)

			// Check the result
			result := a.Elements[:a.Length]
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("InsertOriginal() = %v, want %v", result, tt.expected)
			}
		})
	}
}
