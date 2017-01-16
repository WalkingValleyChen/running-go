package day1

import (
	"fmt"
)

func Iota() {
	const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}

func Function(a, b int) int {
	if (a > b) {
		return a;
	} else {
		return b;
	}

}

func Result2(a, b int) (int, int) {
	return a, b;
}

func FunctionVariable() func(int,int) int{
	mult := 3
	max := func(a, b int) int {
		a *= mult
		b *= mult
		if (a > b) {
			return a;
		} else {
			return b;
		}
	}
	return max;
}
