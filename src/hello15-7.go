package main

import "fmt"

/*
2.1.7 选择器的冷知识

*/

//从一个结构体实例对象中获取字段的值，通常都是使用 . 这个操作符，该操作符叫做 选择器。
//
//选择器有一个妙用，可能大多数人都不清楚。
//
//当你对象是结构体对象的指针时，你想要获取字段属性时，按照常规理解应该这么做

type Profile struct {
	Name string
}

func main() {
	//指针变量
	p1 := &Profile{"iswbm"}
	fmt.Println((*p1).Name) // output: iswbm
	p1.Name = "weibo"
	fmt.Println(p1)
	//&{weibo}

	//但还有一个更简洁的做法，可以直接省去 * 取值的操作，选择器 . 会直接解引用，示例如下

	p2 := &Profile{"cnm"}
	fmt.Println(p2.Name) // output: cnm
}
