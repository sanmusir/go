package main

import (
	"fmt"
)

type TreeNode struct {
	data   interface{}
	Lchild *TreeNode
	Rchild *TreeNode
}

func main() {
	root := CreateTree()
	root.PreOrder()
	fmt.Println()
	root.MidOrder()
	fmt.Println()
	root.PostOrder()
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
	root.Lchild = CreateTree()
	root.Rchild = CreateTree()
	return root
}

//先序遍历
func (root *TreeNode) PreOrder() {
	if root == nil {
		return
	}
	fmt.Print(root.data)
	root.Lchild.PreOrder()
	root.Rchild.PreOrder()
}

//中序遍历
func (root *TreeNode) MidOrder() {
	if root == nil {
		return
	}
	root.Lchild.MidOrder()
	fmt.Print(root.data)
	root.Rchild.MidOrder()
}

//后序遍历
func (root *TreeNode) PostOrder() {
	if root == nil {
		return
	}
	root.Lchild.PostOrder()
	root.Rchild.PostOrder()
	fmt.Print(root.data)
}
