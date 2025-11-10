package main

import (
	"fmt"
	"time"
)

/**
Goroutine
	Go 中的并发执行单位，类似于轻量级的线程。
	Goroutine 的调度由 Go 运行时管理，用户无需手动分配线程。
	使用 go 关键字启动 Goroutine。
	Goroutine 是非阻塞的，可以高效地运行成千上万个 Goroutine。
*/

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)
	}
}

//func main() {
//	go sayHello() // 启动 Goroutine
//	for i := 0; i < 5; i++ {
//		fmt.Println("Main")
//		time.Sleep(100 * time.Millisecond)
//	}
//}

/**
Channel：
	Go 中用于在 Goroutine 之间通信的机制。
	支持同步和数据共享，避免了显式的锁机制。
	使用 chan 关键字创建，通过 <- 操作符发送和接收数据。
*/

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

//func main() {
//	s := []int{7, 2, 8, -9, 4, 0}
//
//	c := make(chan int)
//	go sum(s[:len(s)/2], c)
//	go sum(s[len(s)/2:], c)
//	x, y := <-c, <-c // 从通道 c 中接收
//
//	fmt.Println(x, y, x+y)
//}
