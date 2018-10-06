//go语言算术运算符:+,-,/,*,%,++,--
package main

import "fmt"

func main() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("和：%d\n", c)
	c = a - b
	fmt.Printf("差：%d\n", c)
	c = a * b
	fmt.Printf("积：%d\n", c)
	c = a % b
	fmt.Printf("模：%d\n", c)
	a++
	fmt.Printf("a++ %d\n", a)
	a = 21 //重新为a赋值
	a--
	fmt.Printf("a-- %d\n", a)
}
