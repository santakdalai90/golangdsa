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
