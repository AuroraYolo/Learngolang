package main

//
//import "fmt"
//
///*
//2.1.4 结构体实现"继承"
//为什么标题的继承，加了双引号，因为Go 语言本身并不支持继承。
//但我们可以使用组合的方法，实现类似继承的效果。
//在生活中，组合的例子非常多，比如一台电脑，是由机身外壳，主板，CPU，内存等零部件组合在一起，最后才有了我们用的电脑。
//同样的，在 Go 语言中，把一个结构体嵌入到另一个结构体的方法，称之为组合。
//现在这里有一个表示公司（company）的结构体，还有一个表示公司职员（staff）的结构体。
//*/
//
//type company struct {
//	companyName string
//	companyAddr string
//}
//
//type staff struct {
//	name     string
//	age      int
//	gender   string
//	position string
//	company
//}
//
////若要将公司信息与公司职员关联起来，一般都会想到将 company 结构体的内容照抄到 staff 里。
//
////type staff struct {
////	name        string
////	age         int
////	gender      string
////	companyName string
////	companyAddr string
////	position    string
////}
//
////虽然在实现上并没有什么问题，但在你对同一公司的多个staff初始化的时候，都得重复初始化相同的公司信息，这做得并不好，借鉴继承的思想，我们可以将公司的属性都“继承”过来。
//
////但是在 Go 中没有类的概念，只有组合，你可以将 company 这个 结构体嵌入到 staff 中，做为 staff 的一个匿名字段，staff 就直接拥有了 company 的所有属性了。
//
////type staff struct {
////	name string
////	age int
////	gender string
////	position string
////	company   // 匿名字段
////}
//
//func main() {
//	myCom := company{
//		companyName: "成都",
//		companyAddr: "高新",
//	}
//
//	staffInfo := staff{
//		name:     "小名",
//		age:      28,
//		gender:   "男",
//		position: "程序员",
//		company:  myCom,
//	}
//
//	fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.companyName)
//	fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.company.companyName)
//
//	//输出
//	//小名 在 成都 工作
//	//小名 在 成都 工作
//}
