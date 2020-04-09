package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	//int
	a := 1
	aV(a)
	fmt.Println(a)
	aP(&a)
	fmt.Println(a)
	//string
	s := "测试"
	sV(s)
	fmt.Println(s)
	sP(&s)
	fmt.Println(s)
	//struct
	p := person{
		name: "sanmu",
		age:  18,
	}
	pV(p)
	fmt.Println(p)
	p.edit()
	fmt.Println(p)
	p.update()
	fmt.Println(p)
	//map
	m := map[string] int {
		"a": 1,
		"b": 2,
	}
	mV(m)
	fmt.Println(m)
}

// int 值传递
func aV(a int) {
	a = 2
}

// int 也是值传递，但是传递的是指向a的指针，故a被修改
func aP(a *int) {
	*a = 3
}

// string 值传递
func sV(s string) {
	s = "测试修改"
}

//string 也是值传递，但是传递的是指向s的指针，故s被修改
func sP(s *string) {
	*s = "测试修改"
}

//struct 值传递
func pV(p person) {
	p.age = 20
}

//struct 虽然传递的是指针，但是指针被重新指向了新的结构体，所以原来的结构体并没有被修改
func (p *person) edit() {
	p = &person{
		name: "lin",
		age:  1,
	}
}

//struct 传递的是指针,指针指向的结构体被更新了
func (p *person) update() {
	p.age = 30
	p.name = "lin"
}

//map 声明时返回的是一个指针
func mV(m map[string]int) {
	m["a"] = 2
}

