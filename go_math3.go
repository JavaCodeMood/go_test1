//go语言逻辑运算符：&&,||,!
package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Printf("a && b >>>> true")
	}
	if a || b {
		fmt.Printf("a || b >>>> true\n")
	}
	/* 修改 a 和 b 的值 */
	a = false
	b = true
	if a && b {
		fmt.Printf("a && b >>>> true\n")
	} else {
		fmt.Printf("a && b >>>> false\n")
	}
	if !(a && b) {
		fmt.Printf(" !(a && b) >>>> true\n")
	}
}
