package main

//
//import "fmt"
//
///**
//1.9 流程控制:switch-case
//*/
//func main() {
//
//	/**
//	1.9.0 语句模型
//	switch 表达式 {
//	     case 表达式1:
//	         代码块
//	     case 表达式2:
//	         代码块
//	     case 表达式3:
//	          代码块
//	     default:
//	          代码块
//	}
//	*/
//	//拿switch后的表达式分别和case后的表达式进行对比,只要有一个case满足条件,就会执行对应的代码块,然后直接退出
//	//switch-case,如果一个都没有满足，才会执行default的代码块.
//
//	/**
//	1.9.1 最简单的示例
//	switch后接一个你要判断变量education(学历),然后case会拿这个变量去和它后面的表达式(可能是常量,变量,表达式等)进行判等
//	如果相等,就执行相应的代码块。如果不想等,就接着下一个case
//	*/
//	education := "本科"
//	switch education {
//	case "博士":
//		fmt.Println("我是博士")
//	case "研究生":
//		fmt.Println("我是研究生")
//	case "本科":
//		fmt.Println("我是本科")
//	case "专科":
//		fmt.Println("我是大专生")
//	case "高中":
//		fmt.Println("我是高中生")
//	default:
//		fmt.Println("学历未达标")
//	}
//	//输出如下
//	//我是本科
//
//	/**
//	1.9.2 一个case多个条件
//	case之后可以接多个多个条件,多个条件之间是或的关系,用逗号相隔
//
//	*/
//	month := 2
//	switch month {
//	case 3, 4, 5:
//		fmt.Println("我是春天")
//	case 6, 7, 8:
//		fmt.Println("我是夏天")
//	case 9, 10, 11:
//		fmt.Println("我是秋天")
//	case 12, 1, 2:
//		fmt.Println("我是冬天")
//	default:
//		fmt.Println("输入有误")
//	}
//	//输出如下
//	//我是冬天
//
//	/**
//	1.9.3 case条件不能重复
//	*/
//	//错误案例1
//	//gender := "male"
//	//switch gender {
//	//case "male":
//	//	fmt.Println("男性")
//	//case "male":
//	//	fmt.Println("男性")
//	//case "female":
//	//	fmt.Println("女性")
//	//}
//
//	//错误案例2
//	//gender := "male"
//	//switch gender {
//	//case "male", "male":
//	//	fmt.Println("男性")
//	//case "female":
//	//	fmt.Println("女性")
//	//}
//	/**
//	1.9.4 switch后可接函数
//	switch 后面可以接一个函数,只要保证case后的值类型与函数的返回值一致即可
//	*/
//	//chinese := 80
//	//english := 50
//	//math := 100
//	//switch getResult(chinese, english, math) {
//	//case true:
//	//	fmt.Println("该同学所有成绩都合格")
//	//case false:
//	//	fmt.Println("该同学有挂科记录")
//	//}
//	/**
//	1.9.5 switch 可不接表达式
//	switch 后可以不接任何变量，表达式,函数
//	当不接任何东西时,switch-case 就相当于if-elseif-else
//	*/
//	score := 30
//	switch {
//	case score >= 95 && score <= 100:
//		fmt.Println("优秀")
//	case score >= 80:
//		fmt.Println("良好")
//	case score >= 60:
//		fmt.Println("合格")
//	case score >= 0:
//		fmt.Println("不合格")
//	default:
//		fmt.Println("输入有误")
//	}
//
//	/**
//	1.9.6 switch的穿透能力
//	正常情况下switch-case 的执行顺序是：只要有一个case满足条件，就会直接退出switch-case,如果一个都没有满足,才会执行default的代码块
//	但是有一种情况是例外.
//	那就是当case使用关键字fallthrough开启穿透能力的时候
//	*/
//	//s := "hello"
//	//switch {
//	//case s == "hello":
//	//	fmt.Println("hello")
//	//	fallthrough
//	//case s != "world":
//	//	fmt.Println("word")
//	//}
//	//输出如下
//	//hello
//	//world
//
//	//需要注意的是，fallthrough 只能穿透一层，意思是它让你直接执行下一个case的语句，而且不需要判断条件。
//	s := "hello"
//	switch {
//	case s == "hello":
//		fmt.Println("hello")
//		fallthrough
//	case s == "xxxx":
//		fmt.Println("xxxx")
//	case s != "word":
//		fmt.Println("word")
//	}
//	//输出 并不会输出 world（即使它符合条件）
//	//hello
//	//xxxx
//}
//func getResult(args ...int) bool {
//	for _, i := range args {
//		if i < 60 {
//			return false
//		}
//	}
//	return true
//}
