package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file := os.Stdin
	in := bufio.NewReader(file)
	var T int
	fmt.Fscan(in, &T)
	for i := 1; i <= T; i++ {
		var N int
		fmt.Fscan(in, &N)
		count := 0
		max := -1
		var cur int
		fmt.Fscan(in, &cur)
		for j, next := 1, 0; j < N; j++ {
			fmt.Fscan(in, &next)
			if cur > max {
				max = cur
				if cur > next {
					count++
				}
			}
			cur = next
		}
		if cur > max {
			count++
		}
		fmt.Printf("Case #%d: %d\n", i, count)
	}
}
