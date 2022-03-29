package main

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
	
}
