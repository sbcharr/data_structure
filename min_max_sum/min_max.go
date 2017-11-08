package main

import (
	"fmt"
	"sort"
	"math"
	"os"
)


func main() {
	var list []uint64
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &list[i])
		if list[i] < 1 || list[i] > uint64(math.Pow10(9)) {
			os.Exit(1)
		}
	}
	if !sort.Float64sAreSorted(float64(list)) {
		sort.Float64s(float64(list))
	}
	fmt.Println(list)
	//min := list[0] + list[1] + list[2] + list[3]
	//max := list[4] + list[3] + list[2] + list[1]
	//fmt.Printf("%d %d", min, max)
}


