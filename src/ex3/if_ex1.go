package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)

	score := 100
	switch score {
	case 90, 100:
		fmt.Println("Grade: A")
	case 80:
		fmt.Println("Grade: B")
	case 70:
		fmt.Println("Grade: C")
	case 60:
	case 65:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}

	var a2 = "hello"
	switch a2 {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}

	var a = "mum"
	switch a {
	case "mum", "daddy":
		fmt.Println("family")
	}

	var r int = 11
	switch {
	case r > 10 && r < 20:
		fmt.Println(r)
	}

	sum := 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	// 这样写也可以，更像 While 语句形式
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break
			}
			fmt.Println(i)
		}
	}

	// 遍历, 决定处理第几行
	for y := 1; y <= 9; y++ {
		// 遍历, 决定这一行有多少列
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%d ", x, y, x*y)
		}
		// 手动生成回车
		fmt.Println()
	}

	var str = "hello 你好"
	for key, value := range str {
		fmt.Printf("key:%d value:0x%x\n", key, value)
	}

	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d  value:%d\n", key, value)
	}

	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}

	m2 := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for _, value := range m2 {
		fmt.Println(value)
	}

	c2 := make(chan int)
	go func() {
		c2 <- 1
		c2 <- 2
		c2 <- 3
		close(c2)
	}()
	for v := range c2 {
		fmt.Println(v)
	}

	var breakAgain bool
	// 外循环
	for x := 0; x < 10; x++ {
		// 内循环
		for y := 0; y < 10; y++ {
			// 满足某个条件时, 退出循环
			if y == 2 {
				// 设置退出标记
				breakAgain = true
				// 退出本次循环
				break
			}
		}
		// 根据标记, 还需要退出一次循环
		if breakAgain {
			break
		}
	}
	fmt.Println("done")

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				// 跳转到标签
				goto breakHere
			}
		}
	}
	// 手动返回, 避免执行进入标签
	return
	// 标签
breakHere:
	fmt.Println("breakHere done")

}
