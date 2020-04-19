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
		var N, D int
		fmt.Fscan(in, &N, &D)
		freq := make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Fscan(in, &freq[j])
		}

		for j := N - 1; j >= 0; j-- {
			D = (D / freq[j]) * freq[j]
		}

		fmt.Printf("Case #%d: %d\n", i, D)
	}
}
