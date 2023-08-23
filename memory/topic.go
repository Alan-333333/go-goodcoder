package main

import "fmt"

func foo() {
	x := [1024]int{} // 局部栈数组

	go func(x *[1024]int) {
		// x 发生逃逸
	}(&x)
}

func main() {
	nums := make([]int, 0, 10)
	fmt.Printf("%+v\n", nums)
}
