package main

//
//import "fmt"
//
///*
//2.1.6 三种实例化方法
//
//*/
//
//type Profile struct {
//	name   string
//	age    int
//	gender string
//}
//
//func main() {
//	//第一种: 正常实例化
//	//xm := Profile{
//	//	name:   "小明",
//	//	age:    18,
//	//	gender: "male",
//	//}
//	//fmt.Println(xm)
//
//	//第二种:使用new
//	//xm := new(Profile)
//	// 等价于: var xm *Profile = new(Profile) 指针变量
//	//fmt.Println(xm)
//	//// output: &{ 0 }
//	//
//	//xm.name = "iswbm"
//	//xm.age = 18
//	//xm.gender = "male"
//	//fmt.Println(xm)
//	//var zhizhen = &xm 获取指针内存地址
//
//	//第三种：使用&
//	var xm *Profile = &Profile{}
//	fmt.Println(xm)
//	//&{ 0 }
//	xm.name = "iswbm"  // 或者 (*xm).name = "iswbm"
//	xm.age = 18        //  或者 (*xm).age = 18
//	xm.gender = "male" // 或者 (*xm).gender = "male"
//	fmt.Println(xm)
//	//输出
//    //&{iswbm 18 male}
//}
