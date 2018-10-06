package main

import "fmt"

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
	for i := 1; i < 5; i++ {
		defer fmt.Print(i)
	}
}

func deferIt3() {
	f := func(i int) int {
		fmt.Printf("%d ",i)
		return i * 10
	}
	for i := 1; i < 5; i++ {
		defer fmt.Printf("%d ", f(i))
	}
}

func deferIt4() {
	for i := 1; i < 5; i++ {
		defer func() {
			fmt.Print(i)
		}()
	}
}

func deferIt41() {
	for i := 1; i < 5; i++ {
		defer func(n int) {
			fmt.Print(n)
		}(i)
	}
}


func main(){
	//deferIt()
	deferIt2()
	//deferIt3()
	//deferIt4()
	deferIt41()
}
