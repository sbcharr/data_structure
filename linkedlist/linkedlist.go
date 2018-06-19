package linkedlist

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) AppendNode(val int) {
	node := &Node{
		data: val,
		//next: nil,
	}
	current := l.head
	previous := current
	if current == nil {
		l.head = node
	} else {
		for current != nil {
			previous = current
			current = current.next
		}
		current = node
		previous.next = current
	}

	return
}

func (l *LinkedList) DisplayNode() []int {
	var nodeList []int
	current := l.head
	if l.head != nil {
		for current != nil {
			//fmt.Println(current.data)
			nodeList = append(nodeList, current.data)
			current = current.next
		}
	}
	//fmt.Println("====================")
	return nodeList
}

func (l *LinkedList) UpdateNode(val int, valToReplaceWith int) {
	current := l.head
	if l.head != nil {
		for current != nil {
			if current.data == val {
				current.data = valToReplaceWith
				return
			}
			current = current.next
		}
		fmt.Println("no matching value found")
	}

	return
}

func (l *LinkedList) DeleteNode(val int) {
	previous := l.head
	current := l.head.next

	if l.head != nil && l.head.data == val {
		l.head = current
		previous = nil

		return
	}

	if l.head != nil {
		for current != nil {
			if current.data == val {
				previous.next = current.next
				current = nil
				break
			} else {
				previous = current
				current = current.next
			}
		}
	}
}

func (l *LinkedList) ReverseLinkedList() *Node {
	if l.head == nil || l.head.next == nil {
		fmt.Println("Debug1...")

		return l.head
	}

	current := l.head
	var previous *Node
	todoList := l.head

	for todoList != nil {
		todoList = l.head.next
		l.head = todoList
		current.next = previous
		previous = current
		current = l.head
	}
	l.head = previous
	return l.head
}
