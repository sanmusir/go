package main

import "fmt"
import "time"

func main() {
	f("one")
	go f("two")
	go func(s string) {
		for i := 1; i < 20; i++ {
			fmt.Println(s, ":", i)
		}
	}("three")

	 time.Sleep(2000 * time.Millisecond)
	fmt.Println("done")
	
}

func f(s string) {
	for i := 1; i < 10; i++ {
		fmt.Println(s, ":", i)
	}
}
