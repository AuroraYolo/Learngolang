package main

import "fmt"

/**
数据类型:byte 与 rune与字符串
*/
func main() {
	/**
	1.4.1 byte 与 rune
	//byte 占用1个字节，就8个比特位(2^8=256,因此byte的表示范围0->255),
	 所以它和uint8类型本质上没有区别，它表示的是ASCII表中的一个字符
	*/
	//如下这段代码，分别定义了byte类型和unit8类型的变量a和b
	//var a byte = 65
	//// 8进制写法: var a byte = '\101'     其中 \ 是固定前缀
	//// 16进制写法: var a byte = '\x41'    其中 \x 是固定前缀
	//
	//var b uint8 = 66
	//fmt.Printf("a 的值: %c \nb 的值: %c", a, b)
	//在ASCII表中,由于字母A的ASCII的编号为65,字母B的ASCII编号为66，所以上面的代码也可以写成这样
	//var a byte = 'A'  //65
	//var b uint8 = 'B' //66
	//fmt.Printf("a 的值:%c \nb 的值: %c\n", a, b)
	//输出结果
	//a的值：A
	//b的值: B

	/**
	 rune类型 占用4个字节，共32位比特位，所以它和int32本质上也没有区别。它表示的是一个unicode字符(unicode是一个可以表示世界范围内的绝大
	部分的编码规范)
	*/

	//var a byte = 'A'
	//var b rune = 'B'
	//fmt.Printf("a 占用 %d 个字节数\nb 占用 %d 个字节数", unsafe.Sizeof(a), unsafe.Sizeof(b))
	//输出如下
	//a 占用 1 个字节数
	//b 占用 4 个字节数
	//由于byte类型的表示有限,只有2^8=256个。所以如果你想表示中文的话，你只能使用rune类型
	//var name rune = '中'
	//或许你已经发现，上面我们在定义字符时，不管是 byte 还是 rune ，我都是使用单引号，而没使用双引号。
	//
	//对于从 PHP 转过来的人，这里一定要注意了，在 Go 中单引号与 双引号并不是等价的。
	//
	//单引号用来表示字符，在上面的例子里，如果你使用双引号，就意味着你要定义一个字符串，赋值时与前面声明的会不一致，这样在编译的时候就会出错。
	//cannot use "A" (type string) as type byte in assignment
	//上面我说了，byte 和 uint8 没有区别，rune 和 uint32 没有区别，那为什么还要多出一个 byte 和 rune 类型呢？
	//
	//理由很简单，因为uint8 和 uint32 ，直观上让人以为这是一个数值，但是实际上，它也可以表示一个字符，
	//所以为了消除这种直观错觉，就诞生了 byte 和 rune 这两个别名类型。

	/**
	1.4.2 字符串
	*/

	//var mystr string = "hello"

	//上面说的byte 和 rune 都是字符类型，若多个字符放在一起，就组成了字符串，也就是这里要说的 string 类型。
	//
	//比如 hello ，对照 ascii 编码表，每个字母对应的编号是：104,101,108,108,111

	//var mystr01 string = "hello"
	//var mystr02 [5]byte = [5]byte{104, 101, 108, 108, 111}
	//fmt.Printf("mystr01: %s\n", mystr01)
	//fmt.Printf("mystr02: %s\n", mystr02)
	//输出如下，mystr01和mystr02输出一样，说明了string的本质,其实是一个byte数组
	//mystr01: hello
	//mystr02: hello
	//通过以上学习，我们知道字符分为byte和rune，占用的大小不同

	//"hello,中国"，占用几个字节？Go语言的string是用utf-8进行编码的，英文字母占用一个字节，而中文占用3个字节。所以"hello,中国"的长度
	//为5+1 + (3*2) = 12 个字节
	//var country string = "hello,中国"
	//fmt.Println(len(country))
	//输出 12
	//以上虽然我都用双引号表示 一个字符串，但这并不是字符串的唯一表示方式。
	//
	//除了双引号之外 ，你还可以使用反引号。
	//
	//大多情况下，二者并没有区别，但如果你的字符串中有转义字符\ ，这里就要注意了，它们是有区别的。
	//
	//使用反引号包裹的字符串，相当于 Python 中的 raw 字符串，会忽略里面的转义。
	//
	//比如我想表示 \r\n 这个 字符串，使用双引号是这样写的，这种叫解释型表示法
	//var mystr01 string = "\\r\\n"

	//而使用反引号,就方便多了，所见即所得，这就叫原生型表示法
	//var mystr02 string = `\r\n`

	//fmt.Println(mystr01)
	//fmt.Println(mystr02)

	//输出
	//\r\n
	//\r\n

	//如果你仍然想使用解释型的字符串，但是各种转义实在太麻烦了。你可以使用 fmt 的 %q 来还原一下。
	//var mystr01 string = `\r\n`
	//fmt.Print(`\r\n`)
	//fmt.Printf("的解释型字符串是： %q", mystr01)
	//输出 \r\n的解释型字符串是： "\\r\\n"

	//同时反引号可以不写换行符(没法写)来表示一个多行的字符串。

	var mystr01 = `你好呀！
我的公众号是：Go编程时光，欢迎大家关注
`
	fmt.Println(mystr01)

	//输出
	//你好呀！
	//我的公众号是：Go编程时光，欢迎大家关注
}
