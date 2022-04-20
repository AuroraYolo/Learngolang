package main

import "fmt"

/**
1.8 流程控制:if-else

*/
func main() {
	/**
	1.8.1 条件语句模型
	Go里的流程控制方法还是挺丰富,整理了下有如下这么多种:

	1.if-else条件
	2.switch-case 选择语句
	3.for-range 循环语句
	4.goto无条件跳转语句
	5.defer 延迟执行
	*/

	//今天先来讲讲if-else条件语句
	//Go里的条件语句模型是这样的

	//if 条件1 {
	//分支1
	//} else if 条件2{
	//分支2
	//} else if 条件3 {
	//分支3
	//} else {
	// 分支else
	//}

	//go编译器,对于{和}的位置都有严格的要求，它要求 else if （或else）和两边的花括号，必须在同一行
	//由于go是强类型,所以要求你条件表达式必须严格返回布尔型的数据(nil 和 0和1都不行,具体可查看<<详解数据类型:字典与布尔类型>>

	//对于这个模型,分别举几个列子来看一下.

	/**
	1.8.2单分支判断
	*/
	//只有一个if,没有else

	//age := 20
	//if age > 18 {
	//	fmt.Println("已经成年了")
	//}
	//如何条件里要满足多个条件,可以使用&&和||
	//1. &&:表示且，左右都需要true,最终结果才能为true,否则为false
	//2. ||:表示或,左右只需要一个为true，最终结果为true，否则为false

	//gender := "male"
	//if (age > 18 && gender == "male") {
	//	fmt.Println("是成年男性")
	//}

	/**
	1.8.3 多分支判断
	if-else
	*/
	//age := 20
	//
	//if age > 18 {
	//	fmt.Println("已经是成年了")
	//} else {
	//	fmt.Println("还未成年")
	//}

	/**
	1.8.3 if-else if-else
	*/
	//age := 20
	//if age > 18 {
	//	fmt.Println("已经成年了")
	//} else if age > 12 {
	//	fmt.Println("已经是青年了")
	//} else {
	//	fmt.Println("还不是青少年")
	//}

	/**
	1.8.4 高级写法
	在if里可以写允许先运行一个表达式,取得变量后,在对其进行判断,比如第一个例子里代码也可以写成这样
	*/
	if age := 20; age > 18 {
		fmt.Println("已经成年了")
	}
}
