package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// file := os.Stdin
	in := bufio.NewReader(file)
	var T int
	fmt.Fscan(in, &T)
	for i := 1; i <= T; i++ {
		var R, C int
		fmt.Fscan(in, &R, &C)
		wall := make([][]byte, 0)
		wallMap := make([]map[byte]int, R)
		for j := 0; j < R; j++ {
			var ch byte
			wallMap[j] = make(map[byte]int, 0)
			wall[j] = make([]byte, 0)
			for k := 0; k < C; k++ {
				fmt.Fscan(in, ch)
				wallMap[j][ch] = k
				wall[j][k] = ch
			}
		}

		var solved bool
		solve := make([]byte, 0)
		solvedMap := make(map[byte]bool, 0)

		for j := R - 1; j >= 0; j-- {
			row := wall[j]
			for _, ch := range row {
				
			}
		}

		if solved {
			fmt.Printf("Case #%d: %s\n", i, string(solve))
		} else {
			fmt.Printf("Case #%d: -1\n", i)
		}
	}
}
