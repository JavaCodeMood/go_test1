//go语言别名类型
/*
byte与rune类型有一个共性，即：它们都属于别名类型。byte是uint8的别名类型，而rune则是int32的别名类型。
byte类型的值需用8个比特位表示，其表示法与uint8类型无异。
一个rune类型的值即可表示一个Unicode字符。
在rune类型值的表示中支持几种特殊的字符序列，即：转义符。它们由“\”和一个单个英文字符组成。
*/
package main

import "fmt"

func main() {
	//声明一个rune类型变量并赋值
	var char1 rune = '赞'
	var char2 rune = 'A'

	//字符串格式化函数，%c用于显示rune类型值代表的字符
	fmt.Printf("字符 '%c' 的Unicode代码点是 %s\n", char1, ("U+8D5E"))
	fmt.Printf("字符 '%c' 的Unicode代码点是 %s\n", char2, ("U+0041"))
}
