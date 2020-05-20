package main

import "fmt"

type stack struct {
	data []interface{}
}

type list struct {
	s1 *stack
	s2 *stack
}

func (list *list) push(item interface{}) {
	list.s1.push(item)
}

func (list *list) pop() interface{} {
	if len(list.s2.data) == 0 {
		for len(list.s1.data) != 0 {
			item := list.s1.pop()
			list.s2.push(item)
		}
	}
	if len(list.s2.data) == 0 {
		return nil
	}
	return list.s2.pop()
}

func (stack *stack) push(item interface{}) {
	stack.data = append(stack.data, item)
}

func (stack *stack) pop() interface{} {
	item := stack.data[len(stack.data)-1]
	stack.data = stack.data[:len(stack.data)-1]
	return item
}

func main() {
	s1 := &stack{
		data: []interface{}{1, 2, 3},
	}
	s2 := &stack{
		data: []interface{}{},
	}
	l := &list{
		s1: s1,
		s2: s2,
	}
	l.push(4)
	fmt.Println(l.pop())
	fmt.Println(l.pop())
	fmt.Println(l.pop())
	fmt.Println(l.pop())
	fmt.Println(l.pop())
	l.push(6)
	l.push(7)
	fmt.Println(l.pop())
}
