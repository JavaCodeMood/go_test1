package main

//声明变量，系统自行判断变量类型
var a = "go语言学习"

//指定变量类型，并赋值。如果不赋值，将会使用默认值
var b string = "我爱你啊"

//指定变量类型，不赋值
var c bool

//同时声明多个类型相同的变量，但不赋值
var x, y, z int

//同时声明多个类型相关的变量，并赋值
var e, f = 123, "go"

//这种因式分解关键字的写法一般用于声明全局变量
var (
	age  int = 20
	flag bool
)

func main() {
	/*(这种不带声明格式的只能在函数体中出现)
	省略var,注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误
	*/
	name := "go语言"

	println(a, b, c, name, x, y, z, e, f, age, flag)
}
