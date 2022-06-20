package main

//
//import "fmt"
//
///*
//2.1.2 绑定方法
//*/
//
////在 Go 语言中，我们无法在结构体内定义方法，那如何给一个结构体定义方法呢，答案是可以使用组合函数的方式来定义结构体方法。它和普通函数的定义方式有些不一样，比如下面这个方法
//
///*
//func (person Profile) FmtProfile() {
//	fmt.Printf("名字：%s\n", person.name)
//	fmt.Printf("年龄：%d\n", person.age)
//	fmt.Printf("性别：%s\n", person.gender)
//}
//*/
//
////其中FmtProfile 是方法名，而(person Profile) ：表示将 FmtProfile 方法与 Profile 的实例绑定。我们把 Profile 称为方法的接收者，而 person 表示实例本身，它相当于 Python 中的 self，在方法内可以使用 person.属性名 的方法来访问实例属性。
//// 定义一个名为Profile 的结构体
//
//type Profile struct {
//	name   string
//	age    int
//	gender string
//	mother *Profile // 指针
//	father *Profile // 指针
//}
//
//// 定义一个与 Profile 的绑定的方法
//
//func (person Profile) FmtProfile() {
//	fmt.Printf("名字：%s\n", person.name)
//	fmt.Printf("年龄：%d\n", person.age)
//	fmt.Printf("性别：%s\n", person.gender)
//}
//
//func main() {
//	// 实例化
//	myself := Profile{name: "小明", age: 24, gender: "male"}
//	// 调用函数
//	myself.FmtProfile()
//	//输出如下
//	//名字：小明
//	//年龄：24
//	//性别：male
//}
