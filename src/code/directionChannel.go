package main

import (
	"fmt"
)

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "send a msg")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
