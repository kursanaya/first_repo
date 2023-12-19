package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("1. start run main")
	n := 3
	go multiplyByTwo(n)
	time.Sleep(time.Second) //  suspending
	fmt.Println("4. Hi. Back to main again")
	time.Sleep(time.Second) //  suspending
	fmt.Println("6. done!")
}

func multiplyByTwo(num int) {
	fmt.Println("2. start run go-routine")
	result := num * 2
	fmt.Println("3. result of goroutine: ", result)
	time.Sleep(time.Second)
	fmt.Println("5. back to go-routine again!")
}
