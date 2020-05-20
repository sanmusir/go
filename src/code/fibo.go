package main

import "fmt"

func febo(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	one, two, v := 1, 0, 0
	for i := 2; i <= x; i++ {
		v = one + two
		two = one
		one = v
	}
	return v
}

func main() {
	fmt.Println(febo(6))
}
