package main

import (
	"fmt"
	"strconv"
)

/*
2.2.1 面向对象:接口与多态
在面向对象的领域里，接口一般这样定义：接口定义一个对象的行为。接口只指定了对象应该做什么，至于如何实现这个行为（即实现细节），则由对象本身去确定。

在 Go 语言中，接口就是方法签名（Method Signature）的集合。当一个类型定义了接口中的所有方法，我们称它实现了该接口。这与面向对象编程（OOP）的说法很类似。接口指定
*/

//1.如何定义接口
//使用 type 关键字来定义接口。
//
//如下代码，定义了一个电话接口，接口要求必须实现 call 方法。

//type Phone interface {
//	call()
//}

//2.如何实现接口
//如果有一个类型/结构体，实现了一个接口要求的所有方法，这里 Phone 接口只有 call方法，所以只要实现了 call 方法，我们就可以称它实现了 Phone 接口。
//
//意思是如果有一台机器，可以给别人打电话，那么我们就可以把它叫做电话。
//
//这个接口的实现是隐式的，不像 JAVA 中要用 implements 显示说明。
//
//继续上面电话的例子，我们先定义一个 Nokia 的结构体，而它实现了 call 的方法，所以它也是一台电话。

//type Nokia struct {
//	name string
//}
//
////接收者为Nokia
//func (phone Nokia) call() {
//	fmt.Println("我是Nokia，是一台电话")
//}

//3.接口实现多态
//在 Go 语言中，是通过接口来实现的多态。
//先定义一个商品（Good）的接口，意思是一个类型或者结构体，只要实现了settleAccount() 和 orderInfo() 两个方法，那这个类型/结构体就是一个商品。

type Good interface {
	settleAccount() int
	orderInfo() string
}

//然后我们定义两个结构体，分别是手机和赠品。

type Phone struct {
	name     string
	quantity int
	price    int
}

type FreeGift struct {
	name     string
	quantity int
	price    int
}

//然后分别为他们实现 Good 接口的两个方法

//phone
func (phone Phone) settleAccount() int {
	return phone.quantity * phone.price
}

func (phone Phone) orderInfo() string {
	return "您要购买" + strconv.Itoa(phone.quantity) + "个" +
		phone.name + "计：" + strconv.Itoa(phone.settleAccount()) + "元"
}

// FreeGift
func (gift FreeGift) settleAccount() int {
	return 0
}
func (gift FreeGift) orderInfo() string {
	return "您要购买" + strconv.Itoa(gift.quantity) + "个" +
		gift.name + "计：" + strconv.Itoa(gift.settleAccount()) + "元"
}

//实现了 Good 接口要求的两个方法后，手机和赠品在Go语言看来就都是商品（Good）类型了。
//
//这时候，我挑选了两件商品（实例化），分别是手机和耳机（赠品，不要钱）

//iPhone := Phone{
//	name: "iPhone",
//	quantity:1,
//	price:8000
//}
//earphones := FreeGift{
//name:     "耳机",
//quantity: 1,
//price:    200,
//}
//然后创建一个购物车（也就是类型为 Good的切片），来存放这些商品。
//goods := []Good{iPhone, earphones}

//最后，定义一个方法来计算购物车里的订单金额
func calculateAllPrice(goods []Good) int {
	var allPrice int
	for _, good := range goods {
		fmt.Println(good.orderInfo())
		allPrice += good.settleAccount()
	}
	return allPrice
}
func main() {
	iPhone := Phone{
		name:     "iPhone",
		quantity: 1,
		price:    8000,
	}
	earphones := FreeGift{
		name:     "耳机",
		quantity: 1,
		price:    200,
	}

	goods := []Good{iPhone, earphones}
	allPrice := calculateAllPrice(goods)
	fmt.Printf("该订单总共需要支付 %d 元", allPrice)
	//运行后输出如下:
	//您要购买1个iPhone计：8000元
	//您要购买1个耳机计：0元
	//该订单总共需要支付 8000 元
}
