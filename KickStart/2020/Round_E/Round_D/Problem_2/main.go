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
		var count int
		var pre int
		fmt.Fscan(in, &pre)
		for i, num, upCount, downCount := 1, 0, 0, 0; i < N; i++ {
			fmt.Fscan(in, &num)
			if num > pre {
				upCount++
				downCount = 0
			} else if num < pre {
				downCount++
				upCount = 0
			} else {
				continue
			}
			if upCount == 4 || downCount == 4 {
				count++
				upCount = 0
			}
		}
		fmt.Printf("Case #%d: %d\n", i, count)
	}
}
