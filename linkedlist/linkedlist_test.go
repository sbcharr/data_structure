package linkedlist

import (
	"testing"
)

func TestAppendNode(t *testing.T) {
	var ll = &LinkedList{}
	i := 1
	for ; i <= 10; i++ {
		ll.AppendNode(i)
	}
	nodeList := ll.DisplayNode()
	if len(nodeList) != i-1 {
		t.Errorf("AppendNode function failed: want %d, got %d", i-1, len(nodeList))
	}
}

func TestUpdateNode(t *testing.T) {
	var ll = &LinkedList{}
	i := 1
	for ; i <= 10; i++ {
		ll.AppendNode(i)
	}

	tables := []struct {
		x      int
		y      int
		result int
	}{
		{1, 11, 11},
		{3, 13, 13},
		{10, 20, 20},
	}

	for _, table := range tables {
		ll.UpdateNode(table.x, table.y)
		nodeList := ll.DisplayNode()
		if nodeList[table.x-1] != table.result {
			t.Errorf("UpdateNode function failed: want %d, got %d", table.result, nodeList[table.x-1])
		}
	}
}

func TestDeleteNode(t *testing.T) {
	var ll = &LinkedList{}
	i := 1
	for ; i <= 10; i++ {
		ll.AppendNode(i)
	}
	tables := []struct {
		x      int
		result int
	}{
		{1, 9},
		{3, 8},
		{10, 7},
	}

	for _, table := range tables {
		ll.DeleteNode(table.x)
		nodeList := ll.DisplayNode()
		if len(nodeList) != table.result {
			t.Errorf("DeleteNode function failed: want %d, got %d", table.result, len(nodeList))
		}
	}
}

func TestReverseLinkedListToReturnHead(t *testing.T) {
	var ll = &LinkedList{}
	i := 1
	for ; i <= 10; i++ {
		ll.AppendNode(i)
	}

	n := ll.ReverseLinkedList()

	if n.data != 10 {
		t.Errorf("ReverseLinkedList function failed: want %d, got %d", 10, n)
	}
}


