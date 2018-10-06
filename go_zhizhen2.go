package main

import "fmt"

func main() {
	var a int= 20   /* 声明实际变量 */
	var b *int        /* 声明指针变量 */

	b = &a  /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a  )

	/* 指针变量的存储地址 */
	fmt.Printf("b 变量储存的指针地址: %x\n", b )

	/* 使用指针访问值 */
	fmt.Printf("*b 变量的值: %d\n", *b )
}