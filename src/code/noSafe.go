package main

import (
	"fmt"
	"time"
)

func main() {
	count := 0
	for i := 1; i <= 10000; i++ {
		go func() {
			count++
		}()
	}
	time.Sleep(time.Second * 3)
	fmt.Println(count)
}
