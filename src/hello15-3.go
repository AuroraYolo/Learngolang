package main

import "fmt"

/*
2.1.3 方法的参数传递方式
*/

//当你想要在方法内改变实例的属性的时候，必须使用指针做为方法的接收者。
// 声明一个 Profile 的结构体

type Profile struct {
	name   string
	age    int
	gender string
	mother *Profile // 指针
	father *Profile // 指针
}

// 重点在于这个星号: *
func (person *Profile) increaseAge() {
	person.age += 1
}

func main() {
	myself := Profile{name: "小明", age: 24, gender: "male"}
	fmt.Printf("当前年龄：%d\n", myself.age)
	myself.increaseAge()
	fmt.Printf("当前年龄：%d", myself.age)
	//输出结果 如下，可以看到在方法内部对 age 的修改已经生效。你可以尝试去掉 *，使用值做为方法接收者，看看age是否会发生改变（答案是：不会改变）
	//当前年龄：24
	//当前年龄：25
	fmt.Println(myself.mother)
}

//至此，我们知道了两种定义方法的方式：
//以值做为方法接收者
//以指针做为方法接收者
//那我们如何进行选择呢？以下几种情况，应当直接使用指针做为方法的接收者。
//你需要在方法内部改变结构体内容的时候
//出于性能的问题，当结构体过大的时候
//有些情况下，以值或指针做为接收者都可以，但是考虑到代码一致性，建议都使用指针做为接收者。
//不管你使用哪种方法定义方法，指针实例对象、值实例对象都可以直接调用，而没有什么约束。这一点Go语言做得非常好。
