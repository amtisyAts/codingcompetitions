package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// file, err := os.Open("input.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	file := os.Stdin
	in := bufio.NewReader(file)
	var T int
	fmt.Fscan(in, &T)
	for i := 1; i <= T; i++ {
		var N, K int
		fmt.Fscan(in, &N, &K)
		count := 0
		for j, curVal, num := 0, K, 0; j < N; j++ {
			fmt.Fscan(in, &num)
			if num == curVal {
				if num == 1 {
					count++
					curVal = K
				} else {
					curVal--
				}
			} else {
				if num == K {
					curVal = K - 1
				} else {
					curVal = K
				}
			}
		}
		fmt.Printf("Case #%d: %d\n", i, count)
	}
}
