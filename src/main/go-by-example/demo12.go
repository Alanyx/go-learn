package main

import "fmt"

/**
函数
函数 是 Go 的中心
*/

// 这里是一个函数，接受两个 int 并且以 int 返回它们的和
func plus(a int, b int) int {
	// Go 需要明确的返回值
	// 例如，它不会自动返回最后一个表达式的值
	return a + b
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)
}
