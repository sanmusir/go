package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	producer(ch, &wg)
	wg.Add(1)
	reciver(ch, &wg)
	wg.Wait()
}

func producer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 1; i < 10; i++ {
			fmt.Println(i)
			ch <- i
		}
		fmt.Println("producer done")
		close(ch)
		wg.Done()
	}()
}

func reciver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if v, ok := <-ch; ok {
				fmt.Println(v)
			} else {
				fmt.Println("reciver done")
				break
			}
		}
		wg.Done()
	}()
}
