package main
import (
    "fmt"
)

const (
    KB float64 = 1 << (iota*10)
    MB 
    GB
)


func main() {
    fmt.PrintLn(KB)
    fmt.PrintLn(MB)
    fmt.PrintLn(GB)
}
