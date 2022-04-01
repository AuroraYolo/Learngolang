package main

import "fmt"

/**
数据类型:字典与布尔类型
*/
func main() {
	/**
	1.6.1 字典

	字典（Map 类型），是由若干个 key:value 这样的键值对映射组合在一起的数据结构。

	它是哈希表的一个实现，这就要求它的每个映射里的key，都是唯一的，可以使用 == 和 != 来进行判等操作，换句话说就是key必须是可哈希的。

	什么叫可哈希的？简单来说，一个不可变对象，都可以用一个哈希值来唯一表示，这样的不可变对象，比如字符串类型的对象（可以说除了切片、 字典，函数之外的其他内建类型都算）。

	意思就是，你的 key 不能是切片，不能是字典，不能是函数。。

	字典由key和value组成，它们各自有各自的类型。

	在声明字典时，必须指定好你的key和value是什么类型的，然后使用 map 关键字来告诉Go这是一个字典。
	map[KEY_TYPE]VALUE_TYPE
	*/
	//声明初始化字典

	//第一种方法
	//var scores map[string]int = map[string]int{"english": 80,"chinese": 85}

	//第二种方法
	//scores := map[string]int{"english": 80,"chinese": 85}

	//第三种方法
	//scores := make(map[string]int)
	//scores["english"] = 80
	//scores["chinese"] = 85

	//要注意的是，第一种方法如果拆分称多步(声明，初始化，再赋值),和其他两种有很大的不一样了。相对会毕竟麻烦

	// 声明一个名为 score 的字典
	//var scores map[string]int
	//
	//// 未初始化的 score 的零值为nil，无法直接进行赋值
	//if scores == nil {
	//	// 需要使用 make 函数先对其初始化
	//	scores = make(map[string]int)
	//}
	//
	//scores["chinese"] = 90
	//fmt.Println(scores)
	//
	////字典的相关操作
	//
	////添加元素
	//scores["math"] = 95
	//
	////更新元素,若key已存在，则直接更新value
	//scores["math"] = 100
	//
	////读取元素,直接使用[key]即可,如果key不存在,也不报错,会返回其value-type的零值
	//fmt.Println(scores["math"])
	//
	////删除元素,使用delete函数，如果key不存在,delete函数会静默处理,不会报错.
	//delete(scores, "math")
	//
	//fmt.Println(scores)

	//当访问一个不存在的key时，并不会直接报错,而是会返回这个value的零值,如果value的类型是int，就返回0

	//scores := make(map[string]int)
	//fmt.Println(scores["english"]) //输出0

	//判断key是否存在

	//当key不存在,会返回value-type的零值,所以你不能通过返回的结果是否是零值来判断对应的key是否存在，因为key对应的value值可能恰好就是零值。
	//其实字典的下标读取可以返回两个值,使用第二个返回值都表示对应的key是否存在,若存在ok为true,若不存在,则ok为false

	//scores := map[string]int{"english": 80, "chinese": 85}
	//
	//math, ok := scores["math"]
	//
	//if ok {
	//	fmt.Printf("math 的值是：%d", math)
	//} else {
	//	fmt.Println("math不存在")
	//}

	//如何对字典进行循环
	//Go语言中没有提供类似Python的keys()和values()这样方便的函数，想要获取,你得自己循环.
	//循环还分3种

	//1,获取key和value
	//scores := map[string]int{"english": 80, "chinese": 85}
	//
	//for subject, score := range scores {
	//	fmt.Printf("key: %s, value: %d\n", subject, score)
	//}

	//2.只获取key，这里注意不用占用符
	//scores := map[string]int{"english": 80, "chinese": 85}
	//
	//for subject := range scores {
	//	fmt.Printf("key:%s\n", subject)
	//}

	//3.只获取value,用一个占位符替代

	//scores := map[string]int{"english": 80, "chinese": 85}
	//
	//for _, score := range scores {
	//	fmt.Printf("value:%d\n", score)
	//}

	/**
	1.6.2 布尔类型
	关于布尔值，无非就两个值:true和false.只是这两个值，在不同的语言里可能不同
	在php中，真值用true表示,与1相等,假值使用false表示，与0相等
	而在go中，真值用true表示，不但不与1相等,并且更加严格，不同类型无法进行比较，而假值
	用false表示，同样与0无法比较
	*/

	//var male bool = true
	//fmt.Println(male == 0)
	//Go中确实不如php那样灵活，bool不能与int直接转换，如果要转换，需要自己实现函数

	//go语言对逻辑值取反使用！符号
	//go语言使用&& 表示且,用｜｜表示或。并且有短路行为(即左边表达式已经可以确认整个表达式的值,那么右边将不会被求值)

	var age int = 15
	var gender string = "male"

	//  && 两边的表达式都会执行
	fmt.Println(age > 18 && gender == "male")
	// gender == "male" 并不会执行
	fmt.Println(age < 18 || gender == "male")
}

/**
bool转int
*/
func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

/**
bool转int
*/
func int2bool(b int) bool {
	return b != 0
}
