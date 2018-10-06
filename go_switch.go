//go语言的switch语句
package main

import "fmt"

func main() {
	//定义局部变量
	var grade string = "B"
	var marks int = 99

	switch marks {
	case 100:
		grade = "A"
	case 90:
		grade = "B"
	case 60, 70, 80:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀！\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好！\n")
	case grade == "D":
		fmt.Printf("及格！\n")
	case grade == "F":
		fmt.Printf("不及格！\n")
	default:
		fmt.Printf("叫家长！\n")
	}
	fmt.Printf("你的等级是： %s\n", grade)

	//以下代码用于判断某个 interface 变量中实际存储的变量类型
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型： %T", i)
	case int:
		fmt.Printf("x是int类型")
	case float32:
		fmt.Printf("x是float32类型")
	case float64:
		fmt.Printf("x是float64类型")
	case func(int) float64:
		fmt.Printf("x是func(int)类型")
	case bool, string:
		fmt.Printf("x是bool类型或string类型")
	default:
		fmt.Printf("x是未知类型")
	}

}
