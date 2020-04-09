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

//非递归先序遍历
func (root *TreeNode) PreOrder() {
	t := root
	stack := []*TreeNode{}
	for t != nil || len(stack) != 0 {
		for t != nil {
			fmt.Print(t.data)
			stack = append(stack, t)
			t = t.LChild
		}
		if len(stack) != 0 {
			t = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			t = t.RChild
		}
	}
}

//非递归中序遍历
func (t *TreeNode) MidOrder() {
	stack := []*TreeNode{}
	for t != nil || len(stack) != 0 {
		for t != nil {
			stack = append(stack, t)
			t = t.LChild
		}
		if len(stack) != 0 {
			t = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Print(t.data)
			t = t.RChild
		}
	}
}

//非递归后序遍历
func (t *TreeNode) PostOrder() {
	stack := []*TreeNode{}
	pre := &TreeNode{}
	for t != nil || len(stack) != 0 {
		for t != nil {
			stack = append(stack, t)
			t = t.LChild
		}
		top := stack[len(stack)-1]
		if (top.LChild == nil & top.RChild == nil) || (top.RChild == nil && pre == top.LChild) || pre == top.RChild {
			fmt.Print(top.data)
			pre = top
			stack = stack[:len(stack)-1]
		} else {
			t = top.RChild
		}
	}
}

func main() {
	root := CreateTree()
	root.PreOrder()
	fmt.Println()
	root.MidOrder()
	fmt.Println()
	root.PostOrder()
}
