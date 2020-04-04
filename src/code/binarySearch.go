package main

import (
	"fmt"
)

func main() {
	list := [10]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 100}
	k := 100
	fmt.Println(binarySearch(list, k))
}

func binarySearch(list [10]int, k int) int {
	left := 0
	right := len(list) - 1

	for left <= right {
		mid := (left + right) / 2 
		fmt.Println(left,"---",right,"----",mid,"---",list[mid])
		if list[mid] < k {
			left = mid + 1
		} else if list[mid] > k {
			right = mid - 1
		} else {
			return mid 
		}
	}

	return -1
}
