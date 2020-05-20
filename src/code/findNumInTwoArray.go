package main
import "fmt"

func main() {
	a := [][]int {
		{1,2,8,9},
		{2,4,9,12},
		{4,7,10,13},
		{6,8,11,15},
	}
	find := 3
	row, col := 0,3
	for row < 4 && col >= 0 {
		if a[row][col] == find {
			fmt.Println("yes")
			return
		}
		if a[row][col] > find {
			col--
		}else {
			row++
		}
	}
}


