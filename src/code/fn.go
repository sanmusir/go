ackage main
import (
	"fmt"
)
func main() {
	a := [...]int{3,6,1,5,9,0,4}
	fmt.Print(a)
	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] > a[j] {
				temp := a[j]
				a[j] = a[i]
				a[i] = temp
			}
		}
	}
	fmt.Print(a)
}

