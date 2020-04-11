package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var house, budget int
		fmt.Scan(&house, &budget)
		prices := make([]int, house)
		for j := 0; j < house; j++ {
			fmt.Scan(&prices[j])
		}

		sort.Ints(prices)
		j := 0
		for ; j < house; j++ {
			if prices[j] > budget {
				break
			}
			budget -= prices[j]
		}
		fmt.Printf("Case #%d: %d\n", (i + 1), j)
	}
}
