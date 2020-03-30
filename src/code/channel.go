package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	mc := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			mc <- "send" + strconv.Itoa(i)
		}
	}()

	go func() {
		for  {
			msg := <-mc
			fmt.Println("rec+",msg)
		}
	}()

	time.Sleep(2000 * time.Millisecond)
}
