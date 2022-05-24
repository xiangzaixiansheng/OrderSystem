package main

import (
	"fmt"
	"time"
)

func test_chan() {
	ch1 := make(chan int)
	done := make(chan bool) // 通道
	go func() {
		fmt.Println("子goroutine执行。。。")
		time.Sleep(3 * time.Second)
		data := <-ch1 // 从通道中读取数据
		fmt.Println("data：", data)
		done <- true
	}()

	time.Sleep(5 * time.Second)
	ch1 <- 100
	fmt.Println("等待执行完成***")

	<-done

	fmt.Println("main 函数执行结束")
}
