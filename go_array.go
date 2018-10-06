//go语言定义数组
/*
数组是长度固定，数据类型一致的一个容器
数组的下标从0开始
获取数组的长度：len(数组名)

*/
package main

import "fmt"

func main() {
	//定义一个数组
	var numbers2 [5]int
	numbers2[0] = 2               //2
	numbers2[3] = numbers2[0] - 3 //-1
	numbers2[1] = numbers2[2] + 5 //5
	numbers2[4] = len(numbers2)   //获取数组的长度
	sum := (11)
	// “==”用于两个值的相等性判断
	fmt.Printf("%v\n", (sum == numbers2[0]+numbers2[1]+numbers2[2]+numbers2[3]+numbers2[4]))

	//定义一个字符串数组并赋值
	var arr = [5]string{"abc", "lhf", "bsg", "dhg", "yya"}
	fmt.Printf("字符串数组的长度：%d\n", len(arr))
	fmt.Printf("字符串的第3个值为：%s\n", arr[2])

}
