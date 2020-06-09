package main

import (
	"fmt"
	"sync"
	"time"
)

/** 全局变量 */
var counter int = 0

/** 通过共享内存的方式实现协程间的通信 */
func Memory() {
	start := time.Now()
	lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go func(a, b int) {
			c := a + b

			lock.Lock()
			counter += 1
			fmt.Printf("%d: %d = %d + %d\n", counter, c, a, b)
			lock.Unlock()
		}(1, i)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		//runtime.Gosched()
		if c >= 10 {
			break
		}
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("runtime(s):", consume)
}
