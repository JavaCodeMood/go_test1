//go语言字符串类型
/*
一个字符串类型的值可以代表一个字符序列。这些字符必须是被Unicode编码规范支持的。
字符串的表示法有两种，即：原生表示法和解释型表示法。
若用原生表示法，需用反引号“`”把字符序列包裹起来。
若用解释型表示法，则需用双引号“"”包裹字符序列。
字符串值是不可变的。一旦创建了一个字符串类型的值，就不可能再对它本身做任何修改。
*/
package main

import (
	"fmt"
)

func main() {
	// 声明一个string类型变量并赋值
	var str1 string = "\\\""
	var str2 string = "asjhedjhasdkedj"

	// 这里用到了字符串格式化函数。其中，%q用于显示字符串值的表象值并用双引号包裹。
	fmt.Printf("用解释型字符串表示法表示的 %q 所代表的是 %s。\n", str1, `str1`)
	fmt.Printf("字符串str2的值为：%s\n", str2)
	fmt.Printf("字符串str2的长度为：%d\n", len(str2))
	fmt.Printf("字符串拼接后：%s", str2+str1)

}
