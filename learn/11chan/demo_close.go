package main

import (
	"fmt"
	"time"
)

func test_close() {
	/*
		关闭通道：close(ch)
			子goroutine：写出10个数据
				每写一个，阻塞一次，主goroutine读取一次，解除阻塞

			主goroutine，读取数据
				每次读取数据，阻塞一次，子goroutine，写出一个，解除阻塞

	*/
	ch2 := make(chan int)
	go sendData2(ch2)

	//读取通道的数据
	for {
		time.Sleep(1 * time.Second)
		v, ok := <-ch2 // 最后一次读取
		if !ok {
			fmt.Println("已经读取了所有的数据。。", ok, v)
			break
		}
		fmt.Println("读取的数据：", v, ok)
	}

	fmt.Println("main..over...")

}

func sendData2(ch1 chan int) {
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
}
