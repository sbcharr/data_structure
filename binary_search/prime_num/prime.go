package main

import (
	"fmt"
)

func main() {
	var n int
	prime := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	fmt.Scanf("%d", &n)
	min := 0
	max := len(prime)-1
	var i int
	var flag bool
	for i = 0; i < n; i++ {
		mid := (min + max)/2
		if n == prime[mid] {
			flag = true
			fmt.Println(flag)
			break
		} else if n < prime[mid] {
			max = mid - 1
		} else if n > mid {
			min = mid + 1
		}
	}
	if !flag {
		fmt.Println(flag)
	}
}
