package main
import (
    "fmt"
)

type tz int

func main() {
	a := []string{
		"aa",
		"bb",
	}
	fmt.Println(a)
	for k,_ := range a {
		a[k] = "cc"
	}
	fmt.Println(a)
}


