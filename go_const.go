//go语言常量语法
package main

import "fmt"
import "unsafe"

//常量用做枚举
const (
	name = "刘豆豆"
	age  = 20
	size = len(name)
	num  = unsafe.Sizeof(name)
)

func main() {
	//定义常量
	const HEIGHT int = 10
	const WIDTH int = 5
	//定义局部变量，并指定类型，但是没有赋值
	var area int
	//同时定义多个常量，并赋值
	const a, b, c = 1, false, "str"

	area = HEIGHT * WIDTH
	fmt.Print("面积为 ：", area)
	println() //换行
	println(a, b, c)
	println()
	println(name, age, size, num)

}
