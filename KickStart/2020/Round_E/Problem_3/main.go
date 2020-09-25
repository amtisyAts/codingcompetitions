package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	const SIZE int = 1000000000
	var T int
	fmt.Fscan(in, &T)
	for i := 1; i <= T; i++ {
		var ins string
		fmt.Fscan(in, &ins)
		x, y, _ := getVal(ins)
		x++
		y++
		x %= SIZE
		y %= SIZE
		if x <= 0 {
			x += SIZE
		}
		if y <= 0 {
			y += SIZE
		}
		fmt.Printf("Case #%d: %d %d\n", i, x, y)
	}
}

func getVal(s string) (int, int, int) {
	lenS := len(s)
	var x, y int
	const SIZE int = 1000000000
	for i := 0; i < lenS; i++ {
		switch s[i] {
		case 'N':
			y--
		case 'S':
			y++
		case 'E':
			x++
		case 'W':
			x--
		case ')':
			return x, y, i
		default:
			mul := int(s[i] - '0')
			i += 2
			getX, getY, index := getVal(s[i:])
			i += index
			getX *= mul
			getY *= mul
			x += getX
			y += getY
		}
		x %= SIZE
		y %= SIZE
	}
	return x, y, lenS - 1
}
