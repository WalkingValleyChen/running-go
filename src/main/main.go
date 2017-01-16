package main

import (
	"fmt"
	"day1"
)

func main() {
	fmt.Println("Hello world")
	day1.Iota();
	fmt.Println(day1.Function(1, 2))
	fmt.Println(day1.Result2(2, 3))
	fmt.Println(day1.FunctionVariable()(1,2))
}
