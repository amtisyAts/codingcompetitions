package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var T int
	fmt.Fscan(in, &T)
	for i := 1; i <= T; i++ {
		var W, H, L, U, R, D int
		fmt.Fscan(in, &W, &H, &L, &U, &R, &D)
		// pass, fail := solve(1, 1, W, H, L, U, R, D)
		// result := float64(pass) / (float64(pass) + float64(fail))
		board := make([][]float64, W)
		for i := 0; i < W; i++ {
			board[i] = make([]float64, H)
		}

		board[W-1][H-1] = 1
		for j := H - 2; j >= 0; j-- {
			if chkHole(W-1, j, L, U, R, D) {
				board[W-1][j] = 0
			} else {
				board[W-1][j] = board[W-1][j+1]
			}
		}
		for i := W - 2; i >= 0; i-- {
			board[i][H-1] = board[i+1][H-1]
			for j := H - 2; j >= 0; j-- {
				if chkHole(i, j, L, U, R, D) {
					board[i][j] = 0
				} else {
					board[i][j] = .5*board[i][j+1] + 0.5*board[i+1][j]
				}
			}
		}
		result := strconv.FormatFloat(board[0][0], 'f', -1, 64)
		fmt.Printf("Case #%d: %s\n", i, result)
	}
}

func chkHole(i, j, L, U, R, D int) bool {
	if i >= L-1 && i <= R-1 && j >= U-1 && j <= D-1 {
		return true
	}
	return false
}

// func solve(X, Y, W, H, L, U, R, D int) (int, int) {
// 	if X == W && Y == H {
// 		return 1, 0
// 	} else if X >= L && X <= R && Y >= U && Y <= D {
// 		return 0, 1
// 	}

// 	if X == W {
// 		return solve(X, Y+1, W, H, L, U, R, D)
// 	} else if Y == H {
// 		return solve(X+1, Y, W, H, L, U, R, D)
// 	}
// 	passDown, failDown := solve(X, Y+1, W, H, L, U, R, D)
// 	passRight, failRight := solve(X+1, Y, W, H, L, U, R, D)
// 	return passDown + passRight, failDown + failRight
// }
