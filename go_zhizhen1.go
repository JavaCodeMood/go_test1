package main

import "fmt"

//定义一个接口
type Pet interface {
	Name() string
	Age() uint8
}

//定义结构体
type Dog struct {
	name string
	age  uint8
}

//值方法   实现接口方法
func (dog Dog) Name() string {
	return dog.name
}
func (dog Dog) Age() uint8 {
	return dog.age
}

func main() {
	myDog := Dog{"goggo", 5}
	//使用类型转换表达式把一个Dog类型转换成空接口类型的值
	_, ok1 := interface{}(&myDog).(Pet)
	_, ok2 := interface{}(myDog).(Pet)
	fmt.Printf("%v,%v\n", ok1, ok2)
}
