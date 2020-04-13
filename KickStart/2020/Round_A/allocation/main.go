package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		var house, budget int
		fmt.Scan(&house, &budget)
		prices := make([]int, 1000)
		for j, val := 0, 0; j < house; j++ {
			fmt.Scan(&val)
			prices[val-1]++
		}

		count := 0
		for j := 0; j < 1000; j++ {
			if prices[j] == 0 {
				continue
			}
			cost := prices[j] * (j + 1)
			if budget >= cost {
				budget -= cost
				count += prices[j]
			} else {
				c := budget / (j + 1)
				count += c
				break
			}
		}
		fmt.Printf("Case #%d: %d\n", i, count)
	}
}
