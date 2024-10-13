package linkedlist

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedList(t *testing.T) {
	t.Run("AppendToTail", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.AppendToTail(1)
		list.AppendToTail(2)
		list.AppendToTail(3)

		expected := []int{1, 2, 3}
		result := []int{}
		curr := list.GetHead()
		for curr != nil {
			result = append(result, curr.GetData())
			curr = curr.GetNext()
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("AppendToTail failed. Expected %v, got %v", expected, result)
		}
	})

	t.Run("AppendArray", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.AppendArray([]int{1, 2, 3})

		expected := []int{1, 2, 3}
		result := []int{}
		curr := list.GetHead()
		for curr != nil {
			result = append(result, curr.GetData())
			curr = curr.GetNext()
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("AppendToTail failed. Expected %v, got %v", expected, result)
		}
	})

	t.Run("DeleteNode - delete middle node", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.AppendToTail(1)
		list.AppendToTail(2)
		list.AppendToTail(3)
		list.DeleteNode(2)

		expected := []int{1, 3}
		result := []int{}
		curr := list.GetHead()
		for curr != nil {
			result = append(result, curr.GetData())
			curr = curr.GetNext()
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("DeleteNode failed to delete middle node. Expected %v, got %v", expected, result)
		}
	})

	t.Run("DeleteNode - delete head node", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.AppendToTail(1)
		list.AppendToTail(2)
		list.DeleteNode(1)

		expected := []int{2}
		result := []int{}
		curr := list.GetHead()
		for curr != nil {
			result = append(result, curr.GetData())
			curr = curr.GetNext()
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("DeleteNode failed to delete head node. Expected %v, got %v", expected, result)
		}
	})

	t.Run("DeleteNode - delete tail node", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.AppendToTail(1)
		list.AppendToTail(2)
		list.AppendToTail(3)
		list.DeleteNode(3)

		expected := []int{1, 2}
		result := []int{}
		curr := list.GetHead()
		for curr != nil {
			result = append(result, curr.GetData())
			curr = curr.GetNext()
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("DeleteNode failed to delete tail node. Expected %v, got %v", expected, result)
		}
	})

	t.Run("DeleteNode - delete non-existent node", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.AppendToTail(1)
		list.AppendToTail(2)
		list.AppendToTail(3)
		list.DeleteNode(4)

		expected := []int{1, 2, 3}
		result := []int{}
		curr := list.GetHead()
		for curr != nil {
			result = append(result, curr.GetData())
			curr = curr.GetNext()
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("DeleteNode failed to handle non-existent node. Expected %v, got %v", expected, result)
		}
	})
}

func TestCompareSinglyLinkedList(t *testing.T) {
	tests := []struct {
		name     string
		input1   []int
		input2   []int
		expected bool
	}{
		{"Both Lists Are Empty", []int{}, []int{}, true},
		{"One List Empty, Other Not", []int{}, []int{1, 2, 3}, false},
		{"Other List Empty", []int{1, 2, 3}, []int{}, false},
		{"Equal Lists", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"Lists with Different Elements", []int{1, 2, 3}, []int{1, 3, 2}, false},
		{"Lists with Different Lengths 1", []int{1, 2, 3}, []int{1, 2}, false},
		{"Lists with Different Lengths 2", []int{1, 2}, []int{1, 2, 3}, false},
		{"Single Element Equal", []int{5}, []int{5}, true},
		{"Single Element Not Equal", []int{5}, []int{10}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize first linked list with input1 data
			list1 := NewSinglyLinkedList[int]()
			for _, val := range tt.input1 {
				list1.AppendToTail(val)
			}

			// Initialize second linked list with input2 data
			list2 := NewSinglyLinkedList[int]()
			for _, val := range tt.input2 {
				list2.AppendToTail(val)
			}

			// Compare both lists
			result := CompareSinglyLinkedList(list1, list2)

			// Check if the result matches the expected output
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// TestToArray tests the ToArray function for various cases
func TestToArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty list", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"multiple elements", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"repeated elements", []int{2, 2, 2}, []int{2, 2, 2}},
		{"negative numbers", []int{-5, -10, -15}, []int{-5, -10, -15}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new linked list and append elements
			list := NewSinglyLinkedList[int]()
			list.AppendArray(tt.input)

			// Convert the linked list to array using ToArray method
			result := list.ToArray()

			// Assert the result matches the expected output
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Test function for the String() method of SinglyLinkedList
func TestSinglyLinkedListString(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string
		elements []int
		expected string
	}{
		{
			"EmptyList",
			[]int{},
			"nil",
		},
		{
			"SingleElementList",
			[]int{1},
			"1 -> nil",
		},
		{
			"MultipleElementsList",
			[]int{1, 2, 3, 4},
			"1 -> 2 -> 3 -> 4 -> nil",
		},
		{
			"AllSameElements",
			[]int{5, 5, 5, 5},
			"5 -> 5 -> 5 -> 5 -> nil",
		},
		{
			"LargeValues",
			[]int{100000, 200000, 300000},
			"100000 -> 200000 -> 300000 -> nil",
		},
	}

	// Iterate over each test case and use t.Run
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create linked list and append elements
			list := NewSinglyLinkedList[int]()
			list.AppendArray(tc.elements)

			// Get the string representation of the list
			result := list.String()

			// Compare result with expected
			if result != tc.expected {
				t.Errorf("Test case %s failed: Expected %q, but got %q", tc.name, tc.expected, result)
			}
		})
	}
}
