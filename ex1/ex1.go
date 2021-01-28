// ex1.go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	fmt.Println("My favorite number is", rand.Intn(20))

	num := 10
	for j := 0; j < num; j++ {
		res := getRand(num)
		fmt.Println(res)
	}
	fmt.Println("This is RandSeed:")
	for j := 0; j < num; j++ {
		res := getRandSeed(num)
		fmt.Println(res)
	}
}

func getRand(num int) int {
	var mu sync.Mutex
	mu.Lock()
	v := rand.Intn(num)
	mu.Unlock()
	return v
}

func getRandSeed(num int) int {
	rand.Seed(time.Now().UnixNano())
	var mu sync.Mutex
	mu.Lock()
	v := rand.Intn(num)
	mu.Unlock()
	return v
}
