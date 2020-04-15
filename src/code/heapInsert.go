package main

import (
	"fmt"
	"math"
)

type MinHeap struct {
	data []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		data: []int{math.MinInt64},
	}
}

func (h *MinHeap) Insert(v int) {
	h.data = append(h.data, v)
	i := len(h.data) - 1
	for ; h.data[i/2] > v; i /= 2 {
		h.data[i] = h.data[i/2]
	}
	h.data[i] = v
}

func main() {
	h := NewMinHeap()
	h.Insert(2)
	h.Insert(3)
	h.Insert(4)
	h.Insert(5)
	fmt.Println(h.data)
	h.Insert(1)
	fmt.Println(h.data)
}
