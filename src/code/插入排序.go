package main

import "fmt"

func main() {
	//要排序的数组
	var a = []int{1, 56, 34, 88, 3, 10, 11, 6}
	fmt.Println(a)
	insertSort(a)
	fmt.Println(a)
	insertSortDec(a)
	fmt.Println(a)
}

func insertSort(ap []int) {
	length := len(ap)
	fmt.Println(length)
	//第一个元素不用排序，从第二个开始
	for i := 1; i < length; i++ {
		//取出当前要比较的值
		key := ap[i]
		//取出在他前面的元素,如果比他大就后移一位
		j := i - 1
		for j >= 0 && ap[j] > key {
			ap[j+1] = ap[j]
			j--
		}
		//插入空挡
		ap[j+1] = key
	}
}

func insertSortDec(a []int) {
	length := len(a)
	for i := 1; i < length; i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] < key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
}
