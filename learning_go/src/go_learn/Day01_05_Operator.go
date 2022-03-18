package main

func main() {
	// 简单的算术操作符就略过
	// 1. ^ 异或操作，如果对一个数异或就是按位取反
	// 2. &^ 按位清空, a &^ b 的含义是对于a数，按位遍历b数，如果b数某一位是1，则把a数中对应的那一位置0, 注:
	//    a 和 b 必须是同一种类型的操作数, 不然会报错
	// 3. >> 循环右移
	// 4. << 循环左移
	var a byte = 60  // 00111100
	var b byte = 13  // 00001101
	println(^a)  // 11000011
	println(a&^b)  // 00110000
}