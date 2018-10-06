/*
Go语言-通道的更多种类
    通道有带缓冲和非缓冲之分。缓冲通道中可以缓存N个数据。我们在初始化一个通道值的时候必须指定这个N。
	相对的，非缓冲通道不会缓存任何数据。发送方在向通道值发送数据的时候会立即被阻塞，
	直到有某一个接收方已从该通道值中接收了这条数据。
	非缓冲的通道值的初始化方法如：make(chan int, 0)
	缓冲的通道的初始化方法如：make(chan int, 5)
    注意，非缓冲通道给予make函数的第二个参数值是0。而缓冲通道给予make函数的第二个参数值是一个确定的数值。

    我们还可以以数据在通道中的传输方向为依据来划分通道。默认情况下，通道都是双向的，即双向通道。
	如果数据只能在通道中单向传输，那么该通道就被称作单向通道。
	我们在初始化一个通道值的时候不能指定它为单向。但是，在编写类型声明的时候，我们却是可以这样做的。例如：
type Receiver <-chan int
    类型Receiver代表了一个只可从中接收数据的单向通道类型。这样的通道也被称为接收通道。
	在关键字chan左边的接收操作符<-形象地表示出了数据的流向。
	相对应的，如果我们想声明一个发送通道类型，那么应该这样：
type Sender chan<- int
    这次<-被放在了chan的右边，并且“箭头”直指“通道”。
	我们可以把一个双向通道值赋予上述类型的变量，就像这样：
var myChannel = make(chan int, 3)    //带缓冲的通道
var sender Sender = myChannel    //发送通道
var receiver Receiver = myChannel    //接收通道
    但是，反之则是不行的。像下面这样的代码是通不过编译的：
var myChannel1 chan int = sender
    单向通道的主要作用是约束程序对通道值的使用方式。
	比如，我们调用一个函数时给予它一个发送通道作为参数，以此来约束它只能向该通道发送数据。
	又比如，一个函数将一个接收通道作为结果返回，以此来约束调用该函数的代码只能从这个通道中接收数据。
*/

package main

import (
	"fmt"
	"time"
)

type Sender chan<- int   //发送通道
type Receiver <-chan int //接收通道

func main() {
	//定义一个非缓冲通道
	var myChannel = make(chan int, 0)
	var number int = 6
	//并发执行代码块
	go func() {
		//发送通道
		var sender Sender = myChannel
		sender <- number //向通道sender发送一个数据6
		fmt.Println("发送过去的数据是：", number)
	}()

	go func() {
		//接收通道
		var receiver Receiver = myChannel
		fmt.Println("接收到的数据是：", <-receiver)
	}()
	//让main函数执行结束的时间延迟1秒
	time.Sleep(time.Second)
}
