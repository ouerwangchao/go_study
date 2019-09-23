package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Double(c *gin.Context) {
	go running()

	var input string
	fmt.Scanln(&input)
}

func running() {
	var times int
	for {
		times++
		fmt.Println("tick:", times)
		time.Sleep(time.Second)
	}
}

func Chan(c *gin.Context) {
	fmt.Println(222222)
	fmt.Println(test(33, 33))

	array1 := []string{"a", "b", "c"}
	array2 := modifyArr(array1)
	fmt.Println(array2)

	var chan2 chan int
	fmt.Println(chan2)
	ch1 := make(chan int, 4)
	fmt.Println(ch1)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	for val := range ch1 {
		fmt.Println(val)
	}
	close(ch1)

}

func modifyArr(a []string) []string {
	a[0] = "444"
	a = append(a, "sss")
	return a
}

func test(x, y int) (int, error) {
	return x + y, nil
}

type ss interface {

}
