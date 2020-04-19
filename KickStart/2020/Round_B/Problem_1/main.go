package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var T int
	fmt.Fscan(in, &T)
	for i := 1; i <= T; i++ {
		var N int
		fmt.Fscan(in, &N)
		heights := make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Fscan(in, &heights[j])
		}

		count := 0
		for j := 1; j < N-1; j++ {
			if heights[j] > heights[j-1] && heights[j] > heights[j+1] {
				count++
			}
		}

		fmt.Printf("Case #%d: %d\n", i, count)
	}
}
