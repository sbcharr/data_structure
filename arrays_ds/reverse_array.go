package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var size int
	var array int64
	fmt.Scanf("%d", &size)
	if size < 1 || float64(size) > math.Pow10(3) {
		os.Exit(1)
	}
	ai := make([]int64, size)
	for p := 0; p < size; p++ {
		fmt.Scanf("%d", &array)
		if array < 1 || float64(array) > math.Pow10(4) {
			os.Exit(1)
		}
		ai[p] = array
	}
	for j := size - 1; j >= 0; j-- {
		fmt.Printf("%d ", ai[j])
	}
}
