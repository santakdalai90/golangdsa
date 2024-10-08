package linkedlist

import (
	"reflect"
	"testing"
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
