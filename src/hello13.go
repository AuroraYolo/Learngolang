package main

//
//import (
//	"fmt"
//	"time"
//)
//
///*
//1.14 异常机制:panic 和 recover
//编程语言一般都会有异常捕获机制，在 Python 中 是使用raise 和 try-except 语句来实现的异常抛出和异常捕获的。
//在 Golang 中，有不少常规错误，在编译阶段就能提前告警，比如语法错误或类型错误等，但是有些错误仅能在程序运行后才能发生，比如数组访问越界、空指针引用等，这些运行时错误会引起程序退出。
//当然能触发程序宕机退出的，也可以是我们自己，比如经过检查判断，当前环境无法达到我们程序进行的预期条件时（比如一个服务指定监听端口被其他程序占用），可以手动触发 panic，让程序退出停止运行。
//*/
//
//func set_data(x int) {
//	defer func() {
//		// recover() 可以将捕获到的panic信息打印
//		if err := recover(); err != nil {
//			fmt.Println(err)
//		}
//	}()
//
//	// 故意制造数组越界，触发 panic
//	var arr [10]int
//	arr[x] = 88
//}
//
//func main() {
//	//1。触发panic
//	//手动触发宕机，是非常简单的一件事，只需要调用 panic 这个内置函数即可，就像这样子
//	//panic("crash")
//	//运行后，直接报错宕机
//	//panic: crash
//	//
//	//goroutine 1 [running]:
//	//main.main()
//	//        /Users/heping/GolangWorkSpace/Learngolang/src/hello13.go:12 +0x38
//
//	//2.捕获 panic
//	//发生了异常，有时候就得捕获，就像 Python 中的except 一样，那 Golang 中是如何做到的呢？
//	//
//	//这就不得不引出另外一个内建函数 – recover，它可以让程序在发生宕机后起生回生。
//	//
//	//但是 recover 的使用，有一个条件，就是它必须在 defer 函数中才能生效，其他作用域下，它是不工作的。
//
//	//set_data(20)
//	// 如果能执行到这句，说明panic被捕获了
//	// 后续的程序能继续运行
//	//fmt.Println("everything is ok")
//	//输出
//	//runtime error: index out of range [20] with length 10
//	//everything is ok
//
//	//通常来说，不应该对进入 panic 宕机的程序做任何处理，但有时，需要我们可以从宕机中恢复，至少我们可以在程序崩溃前，做一些操作，举个例子，当 web 服务器遇到不可预料的严重问题时，在崩溃前应该将所有的连接关闭，如果不做任何处理，会使得客户端一直处于等待状态，如果 web 服务器还在开发阶段，服务器甚至可以将异常信息反馈到客户端，帮助调试。
//
//	//3。无法跨协程
//	//从上面的例子，可以看到，即使 panic 会导致整个程序退出，但在退出前，若有 defer 延迟函数，还是得执行完 defer 。
//	//但是这个 defer 在多个协程之间是没有效果，在子协程里触发 panic，只能触发自己协程内的 defer，而不能调用 main 协程里的 defer 函数的。
//	// 这个 defer 并不会执行
//	defer fmt.Println("in main")
//
//	go func() {
//		defer println("in goroutine")
//		panic("")
//	}()
//
//	time.Sleep(2 * time.Second)
//	//输出如下
//	//in goroutine
//	//panic:
//	//
//	//goroutine 4 [running]:
//	//main.main.func1()
//	//        /Users/heping/GolangWorkSpace/Learngolang/src/hello13.go:64 +0x70
//	//created by main.main
//	//        /Users/heping/GolangWorkSpace/Learngolang/src/hello13.go:62 +0x98
//
//	//4.总结一下
//
//	//Golang 异常的抛出与捕获，依赖两个内置函数：
//	//panic：抛出异常，使程序崩溃
//	//recover：捕获异常，恢复程序或做收尾工作
//	//revocer 调用后，抛出的 panic 将会在此处终结，不会再外抛，但是 recover，并不能任意使用，它有强制要求，必须得在 defer 下才能发挥用途。
//
//}
