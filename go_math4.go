//go语言位运算符：&,|,^,<<,>>
package main

import "fmt"

func main() {
	//定义无符号整形
	var a uint = 60 /* 60 = 0011 1100*/
	var b uint = 13 /* 13 = 0000 1101*/
	var c uint = 0

	//&运算：真真为真，有假即假
	c = a & b /*12 = 0000 1100*/
	fmt.Printf("&运算 c 的值为 %d\n", c)

	//|运算：全假为假，有真即真
	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("|运算 c 的值为 %d\n", c)

	//^运算：有真即真，全假全真均为假
	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("^运算 c 的值为 %d\n", c)

	//<<运算：左移运算
	c = a << 2 /* 240 = 1111 0000 */
	fmt.Printf("<<运算 c 的值为 %d\n", c)

	//右移运算
	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Printf(">>运算 - c 的值为 %d\n", c)
}
