package main

import "fmt"
import "errors"

func innerFunc() {
	fmt.Println("Enter innerFunc")
	panic(errors.New("Occur a panic!")) //引发运行时恐慌
	fmt.Println("Quit innerFunc")
}

func outerFunc() {
	fmt.Println("Enter outerFunc")
	innerFunc()
	fmt.Println("Quit outerFunc")
}

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("Fatal error: %s\n", p)
		}
	}()
	fmt.Println("Enter main")
	outerFunc()
	fmt.Println("Quit main")
}

