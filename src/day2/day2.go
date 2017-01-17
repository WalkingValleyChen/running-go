package day2

import (
	"fmt"
	"errors"
)

func Slice() {
	fmt.Println("-----------slice------------")
	arr := []int{1, 2, 3}
	fmt.Println(arr)
	fmt.Println(arr[0:2])
	fmt.Println(arr[2:])
	fmt.Println(arr[:2])

	app := append(arr, 4, 5, 6)
	app1 := make([]int, len(app), cap(app) * 2)
	copy(app1, app)
	fmt.Println(app1)
}

func Range() {
	fmt.Println("-----------range------------")
	arr := []int{1, 2, 3}
	for i, num := range arr {
		fmt.Println("index ", i, "value", num)
	}
}

func Map() {
	fmt.Println("-----------Map------------")
	aMap := map[string]string{"name":"valley", "firstName":"chen"}
	for key, value := range aMap {
		fmt.Println(key, "->", value);
	}
}

func String() {
	fmt.Println("-----------tring------------")
	str := "string"
	c := []byte(str)
	fmt.Println(c)
	fmt.Println(string(c));
	fmt.Println(str[:3])

}

func Error() error {
	fmt.Println("-----------error------------")
	return errors.New("I am a error!")
}

type person struct {
	name string
	age  int
}

func Struct() {
	fmt.Println("-----------Struct------------")

	P := person{"Tom", 25}
	fmt.Println(P.name, P.age)

}
func (p person) print() {
	fmt.Println("name:", p.name, ",age", p.age)
}
func Object() {
	fmt.Println("-----------object------------")
	p := person{"chen", 25}
	p.print();
}

func print() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello world Go！！！", i);
	}

}

func Go() {
	fmt.Println("-----------go------------")
	go print()
}

func num(a int, b int, size int, c chan int) {
	x := a
	y := b;
	for i := 0; i < size; i++ {
		sum:=x+y
		c <-sum
		x=y
		y=sum


	}
}

func Channel() {
	fmt.Println("-----------channel------------")
	c := make(chan int, 10)
	num(1,2,cap(c),c)
	for i:=range c {
		fmt.Println(i)
	}

}