package main

import (
	"fmt"
	// "bugangongwei/HelloWorld/code"
)

func main() {
	// http.InitHttpEngine()

	// fmt.Println(code.SecondMinimum(5, [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}, 3, 5))
	ch := make(chan int, 1)
	ch <- 1
	close(ch)
	value, exist := <-ch
	fmt.Println("channel value:", exist, value)
	value, exist = <-ch
	fmt.Println("channel value:", exist, value)
}
