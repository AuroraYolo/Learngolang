package main

//
//import "fmt"
//
///**
//1.10 流程控制:for 循环
//Go里的流程控制方法还是挺丰富，整理了下有如下这么多种：
//
//if - else 条件语句
//
//switch - case 选择语句
//
//for - range 循环语句
//
//goto 无条件跳转语句
//
//defer 延迟执行
//*/
//func main() {
//	/*
//			0.语句模型
//			for [condition | (int;condition;increment) | Range]
//		    {
//		     statement(s);
//		     }
//		可以看到for后面,可以接三种类型的表达式
//		1.接一个条件表达式
//		2.接三个表达式
//		3.接一个range表达式
//		4.不接表达式
//	*/
//
//	//1.接一个条件表达式
//	a := 1
//	for a <= 5 {
//		fmt.Println(a)
//		a++
//	}
//	/**
//	1
//	2
//	3
//	4
//	5
//	*/
//
//	//2.接三个表达式
//	/**
//	for后面,紧接着三个表达式,使用;分隔.
//	这三个表达式,各有各的用途
//	第一个表达式:初始化控制变量,在整个循环生命周期内，只运行一次
//	第二个表达式:设置循环控制条件,当返回true，继续循环,返回false,结束循环;
//	第三个表达式:每次循环完开始(除第一次)时,给控制变量增量或减量
//	*/
//	for i := 1; i <= 5; i++ {
//		fmt.Println(i)
//	}
//	/**
//	1
//	2
//	3
//	4
//	5
//	*/
//
//	//3 不接表达式:无限循环
//	//for  {
//	//	fmt.Println("z")
//	//}
//	//等价于
//	//for ;;{
//	//	fmt.Println("z222")
//	//}
//
//	var i int = 1
//	for {
//		if i > 5 {
//			break
//		}
//		fmt.Printf("hello,%d\n", i)
//		i++
//	}
//	/*
//		hello,1
//		hello,2
//		hello,3
//		hello,4
//		hello,5
//	*/
//
//	//3 接for-range语句
//	//历一个可迭代对象，是一个很常用的操作。在 Go 可以使用 for-range 的方式来实现。
//	//range 后可接数组、切片，字符串等
//	//由于 range 会返回两个值：索引和数据，若你后面的代码用不到索引，需要使用 _ 表示 。
//	myArr := [...]string{"world", "php", "go"}
//	for _, item := range myArr {
//		fmt.Printf("hello, %s\n", item)
//	}
//	/**
//	hello, world
//	hello, php
//	hello, go
//
//	*/
//	//如果你用一个变量来接收的话，接收到的是索引
//	for i := range myArr {
//		fmt.Printf("hello, %v\n", i)
//	}
//	/*
//		hello, 0
//		hello, 1
//		hello, 2
//	*/
//}
