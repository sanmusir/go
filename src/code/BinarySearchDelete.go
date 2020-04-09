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
		fmt.Print(" ")
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
		if x < t.data {
			t.LChild = t.LChild.Insert(x)
		} else {
			t.RChild = t.RChild.Insert(x)
		}
	}
	return t
}

//找到最小值
func (t *TreeNode) FindMin() interface{} {
	if t == nil {
		return nil
	}
	for {
		if t.LChild == nil {
			return t.data
		}
		t = t.LChild
	}
}

//删除指定元素
func (t *TreeNode) Delete(x int) {
	if t == nil {
		return
	} else if x < t.data {
		t.LChild.Delete(x)
		return
	} else if x > t.data {
		t.RChild.Delete(x)
		return
	} else {
		if t.RChild != nil && t.LChild != nil {
			v := t.LChild.FindMin()
			fmt.Println(v)
			min, _ := v.(int)
			t.data = min
			t.LChild.Delete(min)
			return
		} else if t.RChild == nil && t.LChild != nil {
			*t = *t.LChild
			return
		} else if t.RChild != nil && t.LChild == nil {
			*t = *t.LChild
			return
		}
	}
}

func main() {
	root := &TreeNode{
		data:   6,
		LChild: nil,
		RChild: nil,
	}
	root.Insert(10)
	root.Insert(4)
	root.Insert(20)
	root.Insert(13)
	root.Insert(2)
	root.Insert(5)
	root.Order()
	fmt.Println()
	root.Delete(20)
	root.Order()
}
