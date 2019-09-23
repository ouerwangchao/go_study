package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	//for i := 0; i < 100; i++ {
	//	go echo(i)
	//	echo(i)
	//}
	arr()
}

func echo(i int) {
	fmt.Println(i)
	time.Sleep(20)
}

func arr() {
	var arr []int
	fmt.Println(reflect.TypeOf(arr))
	var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(reflect.TypeOf(balance))
	fmt.Println(balance[0])
	var a []int
	//var b [2]int
	b := make([]int, 10, 23)
	var c int
	var d bool
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(len(a))
	fmt.Println(len(b))
	fmt.Println(cap(b))
	b = append(b,12)
	if(b == nil){
		fmt.Println("2222")
	}
	fmt.Println(b)
}
