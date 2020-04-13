package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Printf("Case #%d: \n", (i + 1))
	}
}
