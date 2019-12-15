package main
import (
    "fmt"
)

func main() {
	a := []int{1,2,3}
	change(a)
	fmt.Println(a)
}

func change(a []int) {
	a[0] = 11
	a[1] = 22
	fmt.Println(a)
}
