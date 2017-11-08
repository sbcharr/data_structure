package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var n, d, element, l, s, p int64
	fmt.Scanf("%d %d", &n, &d)
	if n < 1 || float64(n) > math.Pow10(5) || d < 1 || d > n {
		os.Exit(1)
	}
	ai := make([]int64, n)
	for p = 0; p < n; p++ {
		fmt.Scanf("%d", &element)
		if element < 1 || float64(element) > math.Pow10(6) {
			os.Exit(1)
		}
		ai[p] = element
	}
	/* commented line is O(n2) order, however, implementation below it is the O(n) order of the same problem */
	//for k = 0; k < d; k++ {
	//	v := ai[0]
	//	for l = 0; l < n-1; l++ {
	//		ai[l] = ai[l] + (ai[l+1] - ai[l])
	//	}
	//	ai[n-1] = v
	//}
	newai := make([]int64, n)
	for l = 0; l < n; l++ {
		newai[(l+n-d)%n] = ai[l]
	}
	for s = 0; s < n; s++ {
		fmt.Printf("%d ", newai[s])
	}
}
