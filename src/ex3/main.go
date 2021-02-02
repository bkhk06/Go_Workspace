// ex3 project main.go
package main

import (
	"fmt"
	"os"
	"runtime"

	"./trans"
)

var twoPi = 2 * trans.Pi

func modify(array [5]int) {
	array[0] = 10 // 试图修改数组的第一个元素
	fmt.Println("In modify(), array values:", array)
}
func main() {
	array := [5]int{1, 2, 3, 4, 5} // 定义并初始化一个数组
	modify(array)                  // 传递给一个函数，并试图在函数体内修改这个数组内容
	fmt.Println("In main(), array values:", array)
	/*var a = 6
	var b = 9
	*/
	fmt.Println("\nSum=", Sum(6, 89))

	var (
		HOME   = os.Getenv("HOME")
		USER   = os.Getenv("USER")
		GOROOT = os.Getenv("GOROOT")
	)
	fmt.Println(HOME, USER, GOROOT)

	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
	fmt.Printf("2*Pi = %g\n", twoPi) // 2*Pi = 6.283185307179586
}

func Sum(a, b int) int { return a + b }
