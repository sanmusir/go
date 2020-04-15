package main

import "fmt"

type MinHeap struct {
	data []int
}

func (h *MinHeap) swap(parent int, length int) {
	for {
		child := parent*2 + 1
		//边界
		if child >= length-1 {
			break
		}
		if h.data[child+1] < h.data[child] {
			child++
		}
		if h.data[child] < h.data[parent] {
			h.data[child], h.data[parent] = h.data[parent], h.data[child]
			parent = child
		} else {
			break
		}
	}
}

func main() {
	h := MinHeap{
		data: []int{98, 48, 777, 63, 57, 433, 23, 1112, 1},
	}
	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.swap(i, len(h.data))
	}
	fmt.Println(h.data)
	//排序
	for i := len(h.data) -1;i>=0;i--{
		h.data[0],h.data[i] = h.data[i],h.data[0]
		h.swap(0,i)
	}
	fmt.Println(h.data)


}
