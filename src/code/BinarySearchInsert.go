package main

import "fmt"

type TreeNode struct {
	data   int
	LChild *TreeNode
	RChild *TreeNode
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
		if len(queue) > 1 {
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

//插入数据
func (t *TreeNode) Insert(x int) *TreeNode {
	if t == nil {
		t = &TreeNode{
			data:   x,
			LChild: nil,
			RChild: nil,
		}
	} else {
		fmt.Println(x)
		if x < t.data {
			t.LChild = t.LChild.Insert(x)
		} else {
			t.RChild = t.RChild.Insert(x)
		}
	}
	return t
}

func main() {
	root := &TreeNode{
		data:   5,
		LChild: nil,
		RChild: nil,
	}
	root.Insert(10)
	root.Insert(4)
	fmt.Println()
	root.Order()
}
