package main

import "fmt"

func main() {
	for i := 1; i < 3; i++ {
		defer fmt.Println(i)
	}

	for i := 1; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}
