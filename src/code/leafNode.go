package main

import "fmt"

type TreeNode struct {
	data   interface{}
	LChild *TreeNode
	RChild *TreeNode
}

//遍历叶结点
func (t *TreeNode) LeafNode() {
	if t != nil {
		if t.LChild == nil && t.RChild == nil {
			fmt.Print(t.data)
		}
		if t.LChild != nil {
			t.LChild.LeafNode()
		}
		if t.RChild != nil {
			t.RChild.LeafNode()
		}
	}
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
func main() {
	root := CreateTree()
	root.LeafNode()
	fmt.Println()
}
