package main
import(
	"fmt"
)

func main() {
	s := []int{1,2,3,4,5,6,7,8}
	fmt.Println(len(s),cap(s))
	fmt.Printf("%p \n",&s)
	b := s[:1]
	fmt.Printf("%p \n",&b)
	fmt.Println(len(b),cap(b))
	b = append(b, 11,22,33,44,55,66)
	fmt.Printf("%p \n",&b)
    fmt.Println(len(b),cap(b))
	fmt.Println(b)

	fmt.Println(s)
	fmt.Println(remove(s, 2))
}

func remove(s []int,i int) []int {
	copy(s[i:], s[i+1:])
	fmt.Println(s)
	return s[:len(s)-1] 
}
