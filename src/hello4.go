package main

/**
数据类型:数组与切片
*/
func main() {
	/**
	1.5.1 数组
	数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。因为数组的长度是固定的,所以在
	Go语言中很少直接使用数组.
	*/

	//声明数组,并给该数组里的每个元素赋值(索引值的最小有效值和其他大多数语言一样是0,不是1)

	//var arr [3]int //[3]  里的3 表示该数组的元素个数及容量
	//arr[0] = 1
	//arr[1] = 2
	//arr[2] = 3

	//声明并直接初始化数组

	//第一种方法
	//var arr [3]int = [3]int{1, 2, 3}

	//第二种方法

	//arr := [3]int{1, 2, 3}

	//上面的3表示数组的元素个数，万一你哪天想往该数组中增加元素，你得对应修改这个数字，为了避免这种硬编码，
	//你可以这样写,... 让Go语言自己根据实际情况来分配空间

	//arr := [...]int{1, 2, 3, 4}

	//[3]int 和 [4]int 虽然都是数组，但他们却是不同的类型,使用fmt的%T可以查得。

	//arr01 := [...]int{1, 2, 3}
	//arr02 := [...]int{1, 2, 3, 4}
	//
	//fmt.Printf("%d 的类型是: %T\n", arr01, arr01)
	//fmt.Printf("%d 的类型是: %T", arr02, arr02)
	//输出
	//[1 2 3] 的类型是: [3]int
	//[1 2 3 4] 的类型是: [4]int

	//如果你觉得每次写[3]int 有点麻烦，你可以为[3]int定义一个类型字变量，也就是别名类型。
	//使用type关键字可以定义一个类型字变量，后面只要你想定义一个容器大小为3，元素类型为int的数组
	//都可以使用这个别名类型

	//type arr3 [3]int
	//
	//myarr := arr3{1, 2, 3}
	//
	//fmt.Printf("%d 的类型是: %T", myarr, myarr)
	//输出
	//[1 2 3] 的类型是: main.arr3

	//其实定义数组还有一种偷懒的方法，比如下面这行代码

	//arr := [4]int{2: 3}
	//fmt.Println(arr)

	//输出
	//[0 0 3 0]

	//可以看出[4]int{2:3},4表示数组有4个元素,2和3分别表示该数组索引为2(初始索引为0)的值为3,而其他没有
	//指定值的，就是int类型的零值，即0

	/**
	1.5.2 切片

	切片(Slice)与数组一样,也是可以容纳若干类型相同的元素的容器，与数组不同的是，无法通过切片类型
	来确定某值的长度。每个切片值都会将数组作为其底层层数数据结构。我们也把这样的数组称为切片的底层数组。

	切片是对数组的一个连续片段的引用。所以切片是一个引用类型，这个片段可以是整个数组，也可以是由起始和终止索引标志的一些项的子集。
	需要注意的是，终止索引标志的项不包括在切片内(意思是这是个左闭右开的区间）
	*/

	//myarr := [...]int{1, 2, 3}
	//fmt.Printf("%d 的类型是： %T", myarr[0:2], myarr[0:2])

	//输出
	//[1 2] 的类型是： []int

	//切片的构造，有四种方式

	//1. 对数组进行片段截取，主要有如下两种写法

	//myarr := [5]int{1, 2, 3, 4, 5} //定义一个数组

	//第一种
	//mysli1 := myarr[1:3] // 表示索引从那个1开始，直到到索引为2 (3-1)的元素

	//第二种
	//mysli2 := myarr[1:3:4] //表示索引从1开始，直到索引为2(3-1)的元素

	//在切片时，若不指定第三个数，那么切片终止索引会一直到原数组的最后一个数。而如果指定了第三个数，那么切片终止索引
	//只会到原数组的该索引值

	//myarr := [5]int{1, 2, 3, 4, 5}
	//fmt.Printf("myarr 的长度为：%d，容量为：%d\n", len(myarr), cap(myarr))
	//
	//mysli1 := myarr[1:3]
	//fmt.Printf("mysli1 的长度为：%d，容量为：%d\n", len(mysli1), cap(mysli1))
	//fmt.Println(mysli1)
	//
	//mysli2 := myarr[1:3:4]
	//fmt.Printf("mysli2 的长度为：%d，容量为：%d\n", len(mysli2), cap(mysli2))
	//fmt.Println(mysli2)

	//输出
	//myarr 的长度为：5，容量为：5
	//mysli1 的长度为：2，容量为：4
	//[2 3]
	//mysli2 的长度为：2，容量为：3
	//[2 3]

	//2.从头声明赋值
	//var strList []string
	//
	//var numList []string
	//
	//var numListEmpty []int

	//3.使用make 函数构造，make函数的格式：make([]Type,size,cap)
	//这个函数刚好指出了，一个切片具备的三个要素：类型(Type),长度(size),容量(cap)

	//a := make([]int, 2)
	//b := make([]int, 2, 10)
	//fmt.Println(a, b)
	//fmt.Println(len(a), len(b))
	//fmt.Println(cap(a), cap(b))

	//输出
	//[0 0] [0 0]
	//2 2
	//2 10

	//4.使用和数组一样,偷懒的方法

	//a := []int{4: 2}
	//fmt.Println(a)
	//fmt.Println(len(a), cap(a))

	//输出
	//[0 0 0 0 2]
	//5 5

	//关于len和cap的概念,可能不好理解，这里举个列子:
	//公司名,相当于字面量，也就是变量名
	//公司里的所有工位，相当于已分配到的内存空间
	//公司里的员工，相当于元素
	//cap代表你这个公司最多可以容纳多数员工
	//len代表你这个公司当前有多少个员工

	//由于切片是引用类型，所以你不对它进行赋值的话，它的零值(默认值)是nil。

	//var myarr []int
	//fmt.Println(myarr == nil)
	//
	//数组与切片有相同点，它们都是可以容纳若干类型相同的元素的容器
	//也有不同点,数组的容器大小固定，而切片本身是引用类型，它更像是python中的list，我们可以对它append进行元素的添加

	//myarr := []int{1}
	////追加一个元素
	//myarr = append(myarr, 2)
	////追加多个元素
	//myarr = append(myarr, 3, 4)
	////追加一个切片,...表示解包，不能省略
	//myarr = append(myarr, []int{7, 8}...)
	////在第一个位置插入元素
	//myarr = append([]int{0}, myarr...)
	////在中间插入一个切片(两个元素)
	//myarr = append(myarr[:5], append([]int{5, 6}, myarr[5:]...)...)
	//fmt.Println(myarr)

	//输出
	//[0 1 2 3 4 5 6 7 8]

	//var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//myslice := numbers4[4:6:8]
	//fmt.Printf("myslice为 %d, 其长度为: %d\n", myslice, len(myslice))
	//
	//myslice = myslice[:cap(myslice)]
	//fmt.Printf("myslice的第四个元素为: %d", myslice[3])
}
