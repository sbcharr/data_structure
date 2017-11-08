package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	var n, a, p int64
	var a, ai Interface{}
	fmt.Scanf("%d", &n)
	if n < 2 || float64(n) > math.Pow10(5) {
		os.Exit(1)
	}
	ai := make([]int64, n)
	for p = 0; p < n; p++ {
		fmt.Scanf("%d", &a)
		if float64(a) < (-math.Pow10(9) ) || float64(a) > math.Pow10(9) {
			os.Exit(1)
		}
		ai[p] = a
	}

	sort.Sort(ai)

}
