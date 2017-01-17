package main

import (
	"fmt"
	"day1"
	"day2"
)

func main() {
	fmt.Println("Hello world")
	day1.Iota();
	fmt.Println(day1.Function(1, 2))
	fmt.Println(day1.Result2(2, 3))
	fmt.Println(day1.FunctionVariable()(1,2))
	day1.Array()
	day1.Point()

	day2.Slice()
	day2.Range()
	day2.Map();
	day2.String();
	err:=day2.Error()
	if err!=nil{
		fmt.Println(err)
	}

	day2.Struct()
	day2.Object()
	day2.Go()
	day2.Channel()
}
