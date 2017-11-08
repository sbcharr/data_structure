package main

import (
	"fmt"
	"math"
	"os"
)

type Stack struct {
	Items []int64
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
func (s *Stack) Peek() int64 {
	topItem := s.Items[len(s.Items) - 1]
	return topItem
}

func (s *Stack) Push(item int64) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() {
	if len(s.Items) == 1 {
		s.Items = s.Items[:0]
	} else {
		s.Items = s.Items[0:len(s.Items) - 1]
	}
}

func (s *Stack) MaxInt() (x int64) {
	if s.Size() == 1 {
		x = s.Items[0]
		return
	} else {
		x = s.Items[0]
		for i := 0; i < s.Size() - 1; i++ {
			if x < s.Items[i + 1] {
				x = s.Items[i + 1]
			}
		}
	}
	return
}

func main() {
	s := &Stack{}
	var n, x, k int64
	var query int
	var maxIntArray []int64
	fmt.Scanf("%d", &n)
	if n < 1 || float64(n) > math.Pow10(5) {
		os.Exit(1)
	}
	for k = 0; k < n; k++ {
		fmt.Scanf("%d %d", &query, &x)
		if query < 1 || query > 3 {
			os.Exit(1)
		}

		if query == 1 {
			if x < 1 || float64(x) > math.Pow10(9) {
				fmt.Println("debug1")
				os.Exit(1)
			}
		}
		switch query {
		case 1:
			s.Push(x)
		case 2:
			s.Pop()
		case 3:
			if !s.IsEmpty() {
				maxVal := s.MaxInt()
				maxIntArray = append(maxIntArray, maxVal)
			}
		}
	}
	for p := 0; p < len(maxIntArray); p++ {
		fmt.Println(maxIntArray[p])
	}
}
