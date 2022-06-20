package main

//
//import "fmt"
//
///**
//2.1 面向对象:结构体与继承
//
//*/
//func main() {
//	//0. 什么是结构体？
//	//在之前学过的数据类型中，数组与切片，只能存储同一类型的变量。若要存储多个类型的变量，就需要用到结构体，它是将多个任意类型的变量组合在一起的聚合数据类型。
//	//
//	//每个变量都成为该结构体的成员变量。
//	//
//	//可以理解为 Go语言 的结构体struct和其他语言的class有相等的地位，但是Go语言放弃大量面向对象的特性，所有的Go语言类型除了指针类型外，都可以有自己的方法,提高了可扩展性。
//	//
//	//在 Go 语言中没有没有 class 类的概念，只有 struct 结构体的概念，因此也没有继承，本篇文章，带你学习一下结构体相关的内容。
//
//	//1. 定义结构体
//	//声明结构体
//	//
//	//type 结构体名 struct {
//	//    属性名   属性类型
//	//    属性名   属性类型
//	//    ...
//	//}
//
//	//通过结构体可以定义一个组合字面量，有几个细节，也算是规则，需要你注意。
//
//	//规则一：当最后一个字段和结果不在同一行时，, 不可省略
//
//	xm1 := Profile{
//		name:   "小名",
//		age:    18,
//		gender: "male",
//	}
//
//	///反之，不在同一行，就可以省略。
//
//	xm := Profile1{
//		name:   "小明",
//		age:    18,
//		gender: "male"}
//
//	fmt.Println(xm, xm1)
//
//	//规则二：字段名要嘛全写，要嘛全不写，不能有的写，有的不写。
//	//例如下面这种写法是会报 mixture of field:value and value initializers 错误的
//	//x2 := Profile{
//	//	name: "小明",
//	//	18,
//	//	"male",
//	//}
//
//	//规则三：初始化结构体，并不一定要所有字段都赋值，未被赋值的字段，会自动赋值为其类型的零值。
//
//	x2 := Profile{name: "小明"}
//	fmt.Println(x2.age)
//	//输出0
//
//	//但要注意的是，只有通过指定字段名才可以赋值部分字段。
//	//
//	//若你没有指定字段名，像这样
//	//xm23 := Profile{"小明"}
//	//输出too few values in Profile literal
//
//}
//
////  比如我要定义一个可以存储个人资料名为 Profile 的结构体，可以这么写
//
//type Profile struct {
//	name   string
//	age    int
//	gender string
//	mother *Profile //指针
//	father *Profile //指针
//}
//
////若相邻的属性（字段）是相同类型，可以合并写在一起
//
//type Profile1 struct {
//	name, gender string
//	age          int
//	mother       *Profile // 指针
//	father       *Profile // 指针
//}
