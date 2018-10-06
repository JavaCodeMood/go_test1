/*
Go语言-结构体和方法
    Go语言的结构体类型（Struct）比函数类型更加灵活。它可以封装属性和操作。
	属性即是结构体类型中的字段，而操作则是结构体类型所拥有的方法。
    结构体类型的字面量由关键字type、类型名称、关键字struct，以及由花括号包裹的若干字段声明组成。
	其中，每个字段声明独占一行并由字段名称（可选）和字段类型组成。示例如下：
type Person struct {
    Name   string
    Gender string
    Age    uint8
}
    结构体类型Person中有三个字段，分别是Name、Gender和Age。我们可以用字面量创建出一个该类型的值，像这样：
Person{Name: "Robert", Gender: "Male", Age: 33}
    可以看到，结构体值的字面量（或简称结构体字面量）由其类型的名称和由花括号包裹的若干键值对组成。
	注意，这里的键是其类型中的某个字段的名称（注意，它不是字符串字面量），而对应的值则是欲赋给该字段的那个值。
	另外，如果这里的键值对的顺序与其类型中的字段声明完全相同的话，我们还可以统一省略掉所有字段的名称，就像这样：
Person{"Robert", "Male", 33}
    我们在编写某个结构体类型的值字面量时可以只对它的部分字段赋值，甚至不对它的任何字段赋值。
	这时，未被显式赋值的字段的值则为其类型的零值。注意，这两种情况下，字段的名称是不能被省略的。

   我们在编写一个结构体值的字面量时不需要先拟好其类型。这样的结构体字面量被称为匿名结构体。
   与匿名函数类似，我们在编写匿名结构体的时候需要先写明其类型特征（包含若干字段声明），再写出它的值初始化部分。
   如依据构体类型Person创建一个匿名结构体：
p := struct {
    Name   string
    Gender string
    Age    uint8
}{"Robert", "Male", 33}
    匿名结构体最大的用处就是在内部临时创建一个结构以封装数据，而不必正式为其声明相关规则。
	而在涉及到对外的场景中，我强烈建议使用正式的结构体类型。

  结构体类型可以拥有若干方法（注意，匿名结构体是不可能拥有方法的）。
  所谓方法，其实就是一种特殊的函数。它可以依附于某个自定义类型。方法的特殊在于它的声明包含了一个接收者声明。
  这里的接收者指代它所依附的那个类型。我们仍以结构体类型Person为例。下面是依附于它的一个名为Grow的方法的声明：
func (person *Person) Grow() {
    person.Age++
}
    如上所示，在关键字func和名称Grow之间的那个圆括号及其包含的内容就是接收者声明。
	其中的内容由两部分组成。第一部分是代表它依附的那个类型的值的标识符。第二部分是它依附的那个类型的名称。
	后者表明了依附关系，而前者则使得在该方法中的代码可以使用到该类型的值（也称为当前值）。
	代表当前值的那个标识符可被称为接收者标识符，或简称为接收者。请看下面的示例：
p := Person{"Robert", "Male", 33}
p.Grow()
    我们可以直接在Person类型的变量p之上应用调用表达式来调用它的方法Grow。
	注意，此时方法Grow的接收者标识符person指代的正是变量p的值。这也是“当前值”这个词的由来。
	在Grow方法中，我们通过使用选择表达式选择了当前值的字段Age，并使其自增。
	因此，在语句p.Grow()被执行之后，p所代表的那个人就又年长了一岁（p的Age字段的值已变为34）。

    需要注意的是，在Grow方法的接收者声明中的那个类型是*Person，而不是Person。
	实际上，*Person是Person的指针类型。这也使得person指代的是p的指针，而不是它本身。

    包含若干字段和方法的结构体类型就相当于一个把属性和操作封装在一起的对象。
	结构体类型（以及任何类型）之间都不可能存在继承关系。
	我们可以通过在结构体类型的声明中添加匿名字段（或称嵌入类型）来模仿继承。

    结构体类型属于值类型。它的零值并不是nil，而是其中字段的值均为相应类型的零值的值。
	举个例子，结构体类型Person的零值若用字面量来表示的话则为Person{}。
*/

package main

import "fmt"

//定义结构体
type Person struct {
	Name    string
	Gender  string
	Age     uint8
	Address string
}

//定义一个依附于Person的Move方法
//newaddr 的值是San Francisco
//person.Address的值是Beijing
//这个方法是交换了一下两个地址
func (person *Person) Move(newaddr string) string {
	oldaddr := person.Address
	person.Address = newaddr
	return oldaddr
}

func main() {
	//给结构体赋值
	p := Person{"Robert", "Male", 33, "Beijing"}
	oldAddress := p.Move("San Francisco")
	fmt.Printf("%s moved from %s to %s.\n", p.Name, oldAddress, p.Address)
}
