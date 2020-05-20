package main

import "fmt"

func quickSort(data []int, left int, right int) {
	if left > right {
		return
	}
	i, j := left, right
	base := data[left]

	for i != j {
		for data[j] >= base && i < j {
			j--
		}
		for data[i] <= base && i < j {
			i++
		}
		if i < j {
			data[i], data[j] = data[j], data[i]
		}
	}
	data[left], data[i] = data[i], data[left]
	quickSort(data, left, i-1)
	quickSort(data, i+1, right)
}

func main() {
	a := []int{1, 0, 9, 8, 7, 6, 5, 4, 3, 2}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
