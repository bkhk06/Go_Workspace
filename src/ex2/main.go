// ex2 project main.go
package main

import (
	"fmt"
)

func main() {
	str := "Hello,世界"
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i] // 依据下标取字符串中的字符，类型为byte
		fmt.Println(i, ch)
	}
	fmt.Println("Hello World!")
}
