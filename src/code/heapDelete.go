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

func (h *MinHeap) Delete() (int, error) {
	l := len(h.data)
	if l <= 1 {
		return 0, fmt.Errorf("heap is empty")
	}
	min := h.data[1]
	last := h.data[l-1]
	var i, child int
	for i = 1; i*2 < l; i = child {
		child = i * 2
		//存在右孩子 且 右孩子比左海子小
		if child+1 < l && h.data[child] > h.data[child+1] {
			child++
		}
		if h.data[child] < last {
			h.data[i] = h.data[child]
		} else {
			break
		}
	}
	h.data[i] = last
	h.data = h.data[:l-1]
	return min, nil
}

func main() {
	h := NewMinHeap()
	h.Insert(2)
	h.Insert(3)
	fmt.Println(h.data)
	h.Insert(1)
	fmt.Println(h.data)
	h.Delete()
	fmt.Println(h.data)
	h.Delete()
	fmt.Println(h.data)
	h.Delete()
	fmt.Println(h.data)
	h.Delete()
}
