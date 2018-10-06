/*
Go语言-切片的更多操作方法
在进行“切片”操作的时候需要指定元素下界索引和元素上界索引，就像这样：numbers3[1:4]
我们还可以在方括号中放入第三个正整数，如numbers3[1:4:4]
这第三个正整数被称为容量上界索引。它的意义在于可以把作为结果的切片值的容量设置得更小。
换句话说，它可以限制我们通过这个切片值对其底层数组中的更多元素的访问。
在上一节讲到的numbers3和slice1。针对它们的赋值语句是这样的：
var numbers3 = [5]int{1, 2, 3, 4, 5}
var slice1 = numbers3[1:4]
 这时，变量slice1的值是[]int{2, 3, 4}。
但是我们可以通过如下操作将其长度延展得与其容量相同：
slice1 = slice1[:cap(slice1)]
通过此操作，变量slice1的值变为了[]int{2, 3, 4, 5}，且其长度和容量均为4。
现在，numbers3的值中的索引值在[1,5)范围内的元素都被体现在了slice1的值中。
这是以numbers3的值是slice1的值的底层数组为前提的。
这意味着，我们可以轻而易举地通过切片值访问其底层数组中对应索引值更大的更多元素。
如果我们编写的函数返回了这样一个切片值，那么得到它的程序很可能会通过这种技巧访问到本不应该暴露给它的元素。
这是确确实实是一个安全隐患。

如果我们在切片表达式中加入了第三个索引（即容量上界索引），如：
var slice1 = numbers3[1:4:4]
那么在这之后，无论我们怎样做都无法通过slice1访问到numbers3的值中的第五个元素。
因为这超出了我们刚刚设定的slice1的容量。
如果我们指定的元素上界索引或容量上界索引超出了被操作对象的容量，
那么就会引发一个运行时恐慌（程序异常的一种），而不会有求值结果返回。

虽然切片值在上述方面受到了其容量的限制，但是我们却可以通过另外一种手段对其进行不受任何限制地扩展。
这需要使用到内建函数append。append会对切片值进行扩展并返回一个新的切片值。
使用方法如下：slice1 = append(slice1, 6, 7)
 通过上述操作，slice1的值变为了[]int{2, 3, 4, 6, 7}。
注意，一旦扩展操作超出了被操作的切片值的容量，那么该切片的底层数组就会被自动更换。
这也使得通过设定容量上界索引来对其底层数组进行访问控制的方法更加严谨了。

我们要介绍的最后一种操作切片值的方法是“复制”。
该操作的实施方法是调用copy函数。
该函数接受两个类型相同的切片值作为参数，并会把第二个参数值中的元素复制到第一个参数值中的相应位置（索引值相同）上。
这里有两点需要注意：
  1. 这种复制遵循最小复制原则，即：被复制的元素的个数总是等于长度较短的那个参数值的长度。
  2. 与append函数不同，copy函数会直接对其第一个参数值进行修改。

举例如下：
var slice4 = []int{0, 0, 0, 0, 0, 0, 0}
copy(slice4, slice1)
    通过上述复制操作，slice4会变为[]int{2, 3, 4, 6, 7, 0, 0}。
*/
package main

import "fmt"

func main() {
	//定义一个切片
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//为切片设置容量上界值为8
	slice5 := numbers4[4:6:8] //4是元素下界索引，6是元素上界索引，8是容量上界值
	length := (2)             //5,6   长度
	capacity := (4)           //容量 5,6,7,8
	fmt.Printf("切片slice5长度：%d\n", len(slice5))
	fmt.Printf("切片slice5容量：%d\n", cap(slice5))
	fmt.Printf("%v, %v\n", length == len(slice5), capacity == cap(slice5))
	//将其长度延展得与其容量相同
	slice5 = slice5[:cap(slice5)]

	//切片追加方法
	slice5 = append(slice5, 11, 12, 13) //5,6,7,8,11,12,13
	length = (7)
	fmt.Printf("%v\n", length == len(slice5))
	for i := 0; i < len(slice5); i++ {
		fmt.Printf("切片的第%d个值是：%d\n", i, slice5[i])
	}

	//定义一个切片
	//slice5 := []int[5,6,7,8,11,12,13]
	slice6 := []int{0, 0, 0}

	//切片复制方法 把第二个参数值中的元素复制到第一个参数值中的相应位置上
	copy(slice5, slice6) //0,0,0,8,11,12,13
	e2 := (0)
	e3 := (8)
	e4 := (11)
	fmt.Printf("%v, %v, %v\n", e2 == slice5[2], e3 == slice5[3], e4 == slice5[4])
}