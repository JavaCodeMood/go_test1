//go语言赋值运算符：==,!=,>,<,>=,<=
package main

import "fmt"

func main() {
	var a int = 21
	var b int = 10

	if a == b {
		fmt.Printf("a与b相等")
	} else {
		fmt.Printf("a与b不相等")
	}

	if a < b {
		fmt.Printf("a小于b\n")
	} else {
		fmt.Printf("a大于b\n")
	}

	if a > b {
		fmt.Printf("a大于b\n")
	} else {
		fmt.Printf("a小于 b\n")
	}
	/* Lets change value of a and b */
	a = 5
	b = 20
	if a <= b {
		fmt.Printf("a小于等于b\n")
	}
	if b >= a {
		fmt.Printf("b大于等于a\n")
	}
}
