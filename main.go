package main

import (
	"fmt"

	"www.example.com/algos/heaps"
)

func main() {
	heap := heaps.MaxHeap()
	heap.Inset(10)
	heap.Inset(40)
	heap.Inset(5)
	heap.Inset(16)
	heap.Inset(12)
	fmt.Println(heap.RemoveMax())
	fmt.Println(heap.RemoveMax())
	fmt.Println(heap.RemoveMax())
	fmt.Println(heap.RemoveMax())

}
