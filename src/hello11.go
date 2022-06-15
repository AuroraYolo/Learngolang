package main

import "fmt"

//1.defer 的用法很简单，只要在后面跟一个函数的调用，就能实现将这个 xxx 函数的调用延迟到当前函数执行完后再执行。
//func myfunc() {
//	fmt.Println("B")
//}

var name string = "go"

func myfunc() string {
	defer func() {
		name = "python"
	}()

	fmt.Printf("myfunc 函数里的name：%s\n", name)
	return name
}

func main() {
	//defer myfunc()
	//fmt.Println("a")

	//2.使用 defer 只是延时调用函数，此时传递给函数里的变量，不应该受到后续程序的影响。
	//name := "go"
	//defer fmt.Println(name) // 输出: go
	//
	//name = "python"
	//fmt.Println(name)
	/*
			go
			B
		输出如下，可见给 name 重新赋值为 python，后续调用 defer 的时候，仍然使用未重新赋值的变量值，就好在 defer 这里，给所有的这是做了一个快照一样
	*/

	//如果 defer 后面跟的是匿名函数，情况会有所不同， defer 会取到最后的变量值
	//name := "go"
	//defer func() {
	//	fmt.Println(name) // 输出: python
	//}()
	//name = "python"
	//fmt.Println(name)
	/**
	python
	python
	*/

	//3。多个defer 反序调用
	//name := "go"
	//defer fmt.Println(name) // 输出: go
	//
	//name = "python"
	//defer fmt.Println(name) // 输出: python
	//
	//name = "java"
	//fmt.Println(name)
	/*
		    输出如下，可见 多个defer 是反序调用的，有点类似栈一样，后进先出。
			java
		    python
			go
	*/

	//4.defer与return孰先孰后
	myname := myfunc()
	fmt.Printf("main 函数里的name: %s\n", name)
	fmt.Println("main 函数里的myname: ", myname)
	/*
			myfunc 函数里的name：go
			main 函数里的name: python
			main 函数里的myname:  go
		来一起理解一下这段代码，第一行很直观，name 此时还是全局变量，值还是go

		第二行也不难理解，在 defer 里改变了这个全局变量，此时name的值已经变成了 python

		重点在第三行，为什么输出的是 go ？

		解释只有一个，那就是 defer 是return 后才调用的。所以在执行 defer 前，myname 已经被赋值成 go 了。
	*/

	//4.为什么要有defer
	//看完上面的例子后，不知道你是否和我一样，对这个defer的使用效果感到熟悉？貌似在 Python 也见过类似的用法。
	//虽然 Python 中没有 defer ，但是它有 with 上下文管理器。我们知道在 Python 中可以使用 defer 实现对资源的管理。最常用的例子就是文件的打开关闭。
	//你可能会有疑问，这也没什么意义呀，我把这个放在 defer 执行的函数放在 return 那里执行不就好了。
	//固然可以，但是当一个函数里有多个 return 时，你得多调用好多次这个函数，代码就臃肿起来了。
	//若是没有 defer，你可以写出这样的代码

	/*
		func f() {
		    r := getResource()  //0，获取资源
		    ......
		    if ... {
		        r.release()  //1，释放资源
		        return
		    }
		    ......
		    if ... {
		        r.release()  //2，释放资源
		        return
		    }
		    ......
		    if ... {
		        r.release()  //3，释放资源
		        return
		    }
		    ......
		    r.release()     //4，释放资源
		    return
		}
		使用了 defer 后，代码就显得简单直接，不管你在何处 return，都会执行 defer 后的函数。

		func f() {
		    r := getResource()  //0，获取资源

		    defer r.release()  //1，释放资源
		    ......
		    if ... {
		        ...
		        return
		    }
		    ......
		    if ... {
		        ...
		        return
		    }
		    ......
		    if ... {
		        ...
		        return
		    }
		    ......
		    return
		}
	*/
}
