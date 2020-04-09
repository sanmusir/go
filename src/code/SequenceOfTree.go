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

//层序遍历
func (t *TreeNode) Order() {
	queue := []*TreeNode{}
	if t == nil {
		return
	}
	queue = append(queue, t)
	for len(queue) != 0 {
		t = queue[0]
		fmt.Print(t.data)
		if (len(queue) > 1) {
			queue = queue[1:]
		} else {
			queue = []*TreeNode{}
		}
		if t.LChild != nil {
			queue = append(queue, t.LChild)
		}
		if t.RChild != nil {
			queue = append(queue, t.RChild)
		}
	}
}
func main() {
	root := CreateTree()
	root.Order()
	fmt.Println()
}
