package main

/*
1.11 流程控制:goto 无条件跳转
*/
func main() {
	//0.基本模型
	/*
		goto 后接一个标签，这个标签的意义是告诉 Go程序下一步要执行哪里的代码。

		所以这个标签如何放置，放置在哪里，是 goto 里最需要注意的。

		goto 标签;
		...
		...
		标签: 表达式;
	*/

	//1.最简单的示例
	//	goto flag
	//	fmt.Println("B")
	//flag:
	//	fmt.Println("A")

	//2.如何使用
	//goto 语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能
	//	i := 1
	//
	//flag:
	//	if i <= 5 {
	//		fmt.Println(i)
	//		i++
	//		goto flag
	//	}
	/*
		  1
		2
		3
		4
		5
	*/
	//使用 goto 实现 类型 break 的效果。
	//	i := 1
	//	for {
	//		if i > 5 {
	//			goto flag
	//		}
	//		fmt.Println(i)
	//		i++
	//	}
	//flag:
	/*
		1
		2
		3
		4
		5
	*/

	//使用 goto 实现 类型 continue的效果，打印 1到10 的所有偶数。
	//	i := 1
	//flag:
	//	for i <= 10 {
	//		if i%2 == 1 {
	//			i++
	//			goto flag
	//		}
	//		fmt.Println(i)
	//		i++
	//	}
	/*
		2
		4
		6
		8
		10
	*/

	//3.注意事项
	//goto语句与标签之间不能有变量声明，否则编译错误。

	//	fmt.Println("start")
	//	goto flag
	//	var say = "hello oldboy"
	//	fmt.Println(say)
	//flag:
	//	fmt.Println("end")

	/*
		# go-player
		./hello10.go:85:7: goto flag jumps over declaration of say at ./hello10.go:86:6
	*/
}
