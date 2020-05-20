package main
import "fmt"

func main() {
	a := []int{1,4,5,7,8,2,0}
	len := len(a)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if a[i] > a[j] {
				a[i],a[j] = a[j],a[i]
			}
		}
	}
	fmt.Println(a)
}
