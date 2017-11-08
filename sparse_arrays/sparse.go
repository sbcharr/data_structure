package main

import (
	"fmt"
	"os"
	"math"
)

func main() {
	var n, q int
	var str, query string
	fmt.Scanf("%d", &n)
	if n < 1 || n > 1000 {
		os.Exit(1)
	}
	na := make([]string, n)
	for p := 0; p < n; p++ {
		fmt.Scanf("%s", &str)
		if len(str) < 1 || len(str) > 20 {
			os.Exit(1)
		}
		na[p] = str
	}
	fmt.Scanf("%d", &q)
	if q < 1 || q > 1000 {
		os.Exit(1)
	}
	qa := make([]string, q)
	for o := 0; o < q; o++ {
		fmt.Scanf("%s", &query)
		if len(query) < 1 || len(query) > 20 {
			os.Exit(1)
		}
		qa[o] = query
	}
	var sum, count int
	for j:= 0; j < q; j++ {
		sum = 0
		for i := 0; i < len(na); i++ {
			if na[i] == qa[j] {
				count = 1
			} else {
				count = 0
			}
			sum += count
		}
		fmt.Printf("%d\n", sum)
	}
}

