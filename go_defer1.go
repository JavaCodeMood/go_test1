package main

import "fmt"

//打印斐波那契数列
func fibonacci(num int) int {
	if num == 0 {
		return 0
	}
	if num < 2 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}

//将会逆序输出
func deferIt() {
	defer func() {
		fmt.Print(1)
	}()
	defer func() {
		fmt.Print(2)
	}()
	defer func() {
		fmt.Print(3)
	}()
	fmt.Print(4)
}

func deferIt2() {
	for x := 1; x < 5; x++ {
		defer fmt.Print(x)
	}
}

func deferIt3() {
	f := func(y int) int {
		fmt.Printf("%d ", y)
		return y * 10
	}
	for y := 1; y < 5; y++ {
		defer fmt.Printf("%d ", f(y))
	}
}

func deferIt4() {
	for z := 1; z < 5; z++ {
		defer func() {
			fmt.Print(z)
		}()
	}
}

func deferIt5() {
	for a := 1; a < 5; a++ {
		defer func(m int) {
			fmt.Print(m)
		}(a)
	}
}
func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}

	fmt.Println("\n")
	for k := 0; k < 10; k++ {
		fmt.Printf("%d ", fibonacci(k))
		defer fmt.Printf("%d ", fibonacci(k)) //逆序打印
	}

	fmt.Println("\n")
	for j := 0; j < 10; j++ {
		defer fmt.Printf("%d ", func(n int) int {
			fmt.Printf("%d ", n)
			return n
		}(fibonacci(j)))
	}
	fmt.Println("\n")

	deferIt()
	fmt.Println("\n")
	deferIt2()
	fmt.Println("\n")
	deferIt3()
	fmt.Println("\n")
	deferIt4()
	fmt.Println("\n")
	deferIt5()
}

