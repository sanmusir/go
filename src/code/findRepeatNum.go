package main
import "fmt"

func main() {
	a := []int {2,3,1,0,2,5,3}
	for i := 0; i < len(a); i++ {
loop:	
		for a[i] != i {
			if a[i] == a[a[i]] {
				fmt.Println(a[i])
				break loop
			}
			a[i],a[a[i]] = a[a[i]],a[i]
		}
	}
}
