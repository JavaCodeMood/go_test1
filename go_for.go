/*for循环是一个循环控制结构，可以执行指定次数的循环
Go语言的For循环有3中形式，只有其中的一种使用分号
*/
package main

import "fmt"

func main() {
	var b int = 15
	var a int
	numbers := [6]int{1, 2, 3, 5}

	//1.for循环
	for a := 0; a < 10; a++ {
		fmt.Printf("a1的值为： %d\n", a)
	}

	//2.for循环
	for a < b {
		a++
		fmt.Printf("a2的值为： %d\n", a)
	}

	//3.for循环
	for i, x := range numbers {
		fmt.Printf("第%d位x的值 = %d\n", i, x)
	}

}
