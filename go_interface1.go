//go语言接口
package main

import "fmt"

//定义一个接口
type Animal interface {
	Grow()
	Move(string) string
}

//定义一个结构体类型
type Cat struct {
	Name    string
	Age     int
	Address string
}

//实现接口类型
func (cat *Cat) Grow() {
	fmt.Println("cat grow")
}
func (cat *Cat) Move(address string) string {
	cat.Address = address
	return ""
}
func main() {
	//使用类型转换表达式把一个*Cat类型转换成空接口类型的值
	myCat := Cat{"Little C", 2, "In the house"}
	animal, ok := interface{}(&myCat).(Animal)
	fmt.Printf("%v, %v\n", ok, animal)
}
