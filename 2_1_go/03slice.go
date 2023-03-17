package main

func main() {
	// 因为数组容量是固定的，不能自动扩展。
	// 因为数组作为函数参数是值传递，浪费空间。
	//在go语言中，几乎可以在所有的场景使用切片代替数组。
	//type slice struce {
	//	array *Poniter
	//	len
	//	cap
	//}

	//arr := []int{}

	//数组和切片定义区别
	//创建数组时，[]指定数组长度
	//创建切片时，[]为空，或者使用[...]

	//切片名称 [low: high: max]
	//low: 起始下标位置
	//high: 结束下标位置 len = high - low
	//容量: cap = max -low
	//截取数组时，初始化切片，没有指定切片容量，切片容量跟随原数组

	//append， 追加切片元素，切片的容量会自动增长，1024以下时以2倍方式增长，超过1024时以25%的算法进行增长
	//copy
}
