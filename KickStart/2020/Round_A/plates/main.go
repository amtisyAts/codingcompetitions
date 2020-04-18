package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var T int
	fmt.Fscan(in, &T)
	for i := 1; i <= T; i++ {
		var N, K, P int
		fmt.Fscan(in, &N, &K, &P)
		plates := make([][]int, N)
		h := &PlateHeap{}
		heap.Init(h)
		for n := 0; n < N; n++ {
			plates[n] = make([]int, K)
			for k := 0; k < K; k++ {
				fmt.Fscan(in, &plates[n][k])
			}
			heap.Push(h, [3]int{plates[n][0], n, 0})
		}
		beauty := 0
		for p := 0; p < P; p++ {
			plate := heap.Pop(h).([3]int)
			beauty += plate[0]
			fmt.Println(plate)
			n, k := plate[1], plate[2]
			if k != (K - 1) {
				heap.Push(h, [3]int{plates[n][k+1], n, k + 1})
			}
		}
		fmt.Printf("Case #%d: %d\n", i, beauty)
	}
}

// An PlateHeap is a max-heap of ints.
type PlateHeap [][3]int

func (h PlateHeap) Len() int           { return len(h) }
func (h PlateHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h PlateHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds a new data to heap
func (h *PlateHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([3]int))
}

// Pop removes the maximum element from heap
func (h *PlateHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
