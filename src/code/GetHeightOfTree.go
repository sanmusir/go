package main

import "fmt"

type TreeNode struct {
	data   interface{}
	LChild *TreeNode
	RChild *TreeNode
}

//创建结点
func CreateTree() *TreeNode {
	num := -1
	fmt.Scan(&num)
	if num == -1 {
		return nil
	}

	root := &TreeNode{
		data: num,
	}
	root.LChild = CreateTree()
	root.RChild = CreateTree()
	return root
}

func (t *TreeNode) Height() int {
	if t == nil {
		return 0
	}
	MaxH, LH, RH := 0, 0, 0
	LH = t.LChild.Height()
	RH = t.RChild.Height()
	if LH > RH {
		MaxH = LH
	} else {
		MaxH = RH
	}
	return MaxH + 1
}

func main() {
	root := CreateTree()
	height := root.Height()
	fmt.Println(height)
}
