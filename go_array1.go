//go语言数组
package main

import "fmt"

func main() {

	var arr1 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//遍历数组
	for a := 0; a < len(arr1); a++ {
		fmt.Printf("数组的第 %d 个值是：%d\n", a, arr1[a])
	}

}
