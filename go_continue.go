//gogo语言循环控制语句---continue语句
package main

import "fmt"

func main() {
	//定义局部变量
	var a int = 10

	//for循环
	for a < 20 {
		if a == 15 {
			a = a + 1
			continue //跳过本次循环
		}
		fmt.Printf("a 的值为：%d\n", a)
		a++
	}
}
