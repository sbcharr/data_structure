package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	Items []string
}

func (s *Stack) IsEmpty() bool {
	if len(s.Items) == 0 {
		return true
	}
	return false
}


//Returns top item from the stack
func (s *Stack) Peek() string {
	topItem := s.Items[len(s.Items) - 1]
	return topItem
}


func (s *Stack) Push(item string) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() string {
	var copyInto []string
	var item string
	if len(s.Items) == 1 {
		item = s.Peek()
		s.Items = s.Items[:0]
		return item
	} else if len(s.Items) == 0 {
		return ""
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
	var n int
	var str string
	fmt.Scanf("%d", &n)
	tokens := [][]string{
		{"{", "}"},
		{"[", "]"},
		{"(", ")"},
	}
	balanced := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &str)
		st := &Stack{}
		var flag bool
		for _, c := range str {
			if st.IsEmpty() {
				if !strings.Contains("{[(", string(c)) {
					flag = true
					balanced[i] = "NO"
					//fmt.Println(balanced[i])
					break
				}
			} else {
				if strings.Contains("}])", string(c)) {
					topItem := st.Peek()

					for l, _ := range tokens {
						if tokens[l][1] == string(c) && topItem == tokens[l][0] {
							st.Pop()
							break
						}
					}
					continue
				}
			}
			st.Push(string(c))
		}
		if st.IsEmpty() && !flag {
			balanced[i] = "YES"
		} else {
			balanced[i] = "NO"
		}
	}
	for _, val := range balanced {
		fmt.Println(val)
	}
}

