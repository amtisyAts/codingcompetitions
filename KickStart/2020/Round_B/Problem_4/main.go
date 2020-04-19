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
		var W, H, L, U, R, D int
		fmt.Fscan(in, &W, &H, &L, &U, &R, &D)
		pass, fail := solve(1, 1, W, H, L, U, R, D)
		result := float64(pass) / (float64(pass) + float64(fail))
		fmt.Printf("Case #%d: %f\n", i, result)
	}
}

func solve(X, Y, W, H, L, U, R, D int) (int, int) {
	if X == W && Y == H {
		return 1, 0
	} else if X >= L && X <= R && Y >= U && Y <= D {
		return 0, 1
	}

	if X == W {
		return solve(X, Y+1, W, H, L, U, R, D)
	} else if Y == H {
		return solve(X+1, Y, W, H, L, U, R, D)
	}
	passDown, failDown := solve(X, Y+1, W, H, L, U, R, D)
	passRight, failRight := solve(X+1, Y, W, H, L, U, R, D)
	return passDown + passRight, failDown + failRight
}
