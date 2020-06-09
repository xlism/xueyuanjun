package main

import (
	"fmt"
	"time"
)

func add(a int, b int, ch chan int) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	ch <- 1
}

func Timeout() {
	ch := make(chan int, 1)

	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()

	select {
		case <-ch:
			fmt.Println("receive the ch channel of data")
		case <-timeout:
			fmt.Println("timeout one seconds, exit")
	}
}

func Channel6(){
	ch := make(chan int, 2)

	go func() {
		for i:=0; i<5; i++ {
			fmt.Printf("发送方：发送数据 %v...\n", i)
			ch<-i
		}
		fmt.Println("发送方：关闭通道...")
		close(ch)
	}()

	for  {
		num, ok := <- ch
		if !ok {
			fmt.Println("接收方：通道已关闭")
			break
		}
		fmt.Printf("接收方：接收数据 %v...\n", num)
	}
	fmt.Println("程序退出")
}
