package main

import (
	"fmt"
	"math"
	"os"
	//"time"
)

type List struct {
	elements []int64
}

func (l *List) AddElements(a, b, k int64) {
	for i := a-1; i < b; i++ {
		l.elements[i] += k
	}

}
/*
func main() {
	var size, ops, i, a, b, k int64
	fmt.Scanf("%d %d", &size, &ops)
	switch {
	case size < 3 || float64(size) > math.Pow10(7): os.Exit(1)
	case ops < 1 || float64(ops) > 2 * math.Pow10(5): os.Exit(1)
	}
	list := &List{}
	list.elements = make([]int64, size)
	//t := time.Now()
	for i = 0; i < ops; i++ {
		fmt.Scanf("%d %d %d", &a, &b, &k)
		validate(a, b, k, size)
		//list.AddElements(a, b, k)
	}
	//fmt.Println(maxInList(list.elements))
	//t2 := time.Since(t)
	//fmt.Println(t2)

}
*/
func validate(a, b, k, size int64) {
	switch {
	case a < 1 || a > b || a > size: os.Exit(1)
	case b > size: os.Exit(1)
	case k < 0 || float64(k) > math.Pow10(9): os.Exit(1)
	}
}

func maxInList(slice []int64) (maxVal int64) {
	var j int64
	maxVal = slice[0]
	for j = 0; j < int64(len(slice)-1); j++ {
		if maxVal < slice[j+1] {
			maxVal = slice[j+1]
		}
	}
	return
}
/*
func AddElements2(a, b, k int64) {
	//el := make([]int64, 10000000)
	for i := a-1; i < b; i++ {
		el[i] += k
	}

}
*/
func main() {
	el := make([]int64, 10000000)
	var a, b, k int64
	a = 10
	b = 20
	k = 5
	el[a:b] = el[a:b] + k
	fmt.Println(el[a:b])
}



