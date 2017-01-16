package day1

import (
	"fmt"
)

/*
好神奇的iota
 */
func Iota() {
	fmt.Println("-----------iota------------")
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
	fmt.Println("-----------func------------")
	if (a > b) {
		return a;
	} else {
		return b;
	}

}

/*
多返回值
 */
func Result2(a, b int) (int, int) {
	return a, b;
}

/*
有闭包功能
 */
func FunctionVariable() func(int, int) int {
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

func Array() {
	fmt.Println("-----------array------------")
	arr := []int{1, 2, 3}
	for i := 0; i < ArrayParam(arr); i++ {
		arr[i] = arr[i] * 2
	}
	fmt.Println(arr)
}
/*
好像不需要传递size
*/
func ArrayParam(arr[]int) int {
	return len(arr)
}

/*
指针
 */
func Point(){
	fmt.Println("-----------point------------")
	var ip *int;
	i:=10;
	fmt.Println("point is nil:",ip==nil)
	ip=&i
	fmt.Println(&i)
	fmt.Println(ip)
	fmt.Println(*ip)
}