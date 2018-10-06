//go语言的可变常量-iota
package main

import "fmt"

//全局常量枚举
const (
	x = 1 << iota
	y = 3 << iota
	z
	w = 3 >> iota
)

func main() {
	//局部常量枚举
	const (
		a = iota  //0
		b         //1
		c         //2
		d = "lhf" //独立值，iota += 1
		e         //"lhf" iota += 1
		f = 100   //iota += 1
		g         //100 iota += 1
		h = iota  //7,恢复计数
		i         //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
	println()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
	fmt.Println("z=", z)
	fmt.Println("w=", w)
}
