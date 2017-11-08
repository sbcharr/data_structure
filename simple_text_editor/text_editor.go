package main

import (
	"fmt"
)

type Stack struct {
	Items []interface{}
}

func (s *Stack) IsEmpty() bool {
	if len(s.Items) == 0 {
		return true
	}
	return false
}

func (s *Stack) Size() int {
	size := len(s.Items)
	return size
}

//Returns top item from the stack
func (s *Stack) Peek() interface{} {
	topItem := s.Items[len(s.Items) - 1]
	return topItem
}

func (s *Stack) Push(item interface{}) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() interface{} {
	var copyInto []interface{}
	var item interface{}
	if len(s.Items) == 1 {
		item = s.Peek()
		s.Items = s.Items[:0]
		return item
	} else if len(s.Items) == 0 {
		return nil
	} else {
		for i := 0; i < len(s.Items) - 1; i++ {
			copyInto = append(copyInto, s.Items[i])
		}
		item = s.Peek()
		s.Items = copyInto
	}
	return item
}

func main() {
	s := &Stack{}

}

