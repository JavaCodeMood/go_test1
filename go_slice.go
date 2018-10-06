//go语言切片类型
/*
切片（Slice）与数组不同的是，无法通过切片类型来确定其值的长度。
每个切片值都会将数组作为其底层数据结构。我们也把这样的数组称为切片的底层数组。
切片类型的字面量如：[]int 或[]string
切片表达式一般由字符串、数组或切片的值以及由方括号包裹由英文冒号“:”分隔的两个正整数组成。
这两个正整数分别表示元素下界和上界索引。
一个切片值的容量即为它的第一个元素值在其底层数组中的索引值与该数组长度的差值的绝对值。
切片类型属于引用类型。它的零值为nil，即空值。
*/
package main

import "fmt"

func main() {
	//定义一个数组
	var numbers3 = [5]int{1, 2, 3, 4, 5}
	//定义一个切片
	slice3 := numbers3[2:len(numbers3)]
	length := (3)
	capacity := (3)
	fmt.Printf("%v, %v\n", (length == len(slice3)), (capacity == cap(slice3)))
	fmt.Printf("切片slice3的容量为：%d\n", cap(slice3))
	fmt.Printf("数组numbers3的长度为：%d\n", len(numbers3))
}
