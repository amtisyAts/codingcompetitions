package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	scanner := getScanner(os.Stdin)
	// Sets words as the split function. i.e. separated by space
	scanner.Split(bufio.ScanWords)
	N, _ := readInt(scanner)
	for i, house, budget := 1, 0, 0; i <= N; i++ {
		house, _ = readInt(scanner)
		budget, _ = readInt(scanner)
		prices := make([]int, 1000)
		for j, val := 0, 0; j < house; j++ {
			val, _ = readInt(scanner)
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

// getScanner returns a scanner buffer from input file/io e.g. os.StdIn
func getScanner(input *os.File) *bufio.Scanner {
	in := bufio.NewReader(input)
	scanner := bufio.NewScanner(in)
	return scanner
}

// readInt returns the next integer from with scanner input. It also returns any
// error encountered
func readInt(scanner *bufio.Scanner) (int, error) {
	// scan for the next token
	if !scanner.Scan() {
		return 0, io.ErrUnexpectedEOF
	}
	// manually parse existing bytes without allocation
	i, err := atoi(scanner.Bytes())
	return i, err
}

// atoi parses a buffer into an integer
// without overflow detection.
func atoi(b []byte) (int, error) {
	neg := false
	if b[0] == '+' {
		b = b[1:]
	} else if b[0] == '-' {
		neg = true
		b = b[1:]
	}
	n := 0
	for _, v := range b {
		if v < '0' || v > '9' {
			return 0, strconv.ErrSyntax
		}
		n = n*10 + int(v-'0')
	}
	if neg {
		return -n, nil
	}
	return n, nil
}
