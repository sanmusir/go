package main
import(
    "fmt"
)

type newInt int
const (
    A newInt = 1 << iota 
    B
    C
)

func main() {
    d := A | B
    fmt.Print(A,B,C)
    fmt.Print(d)
}
