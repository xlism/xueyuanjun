package main

import (
	"fmt"
	"sync"
)

func Pool() {
	pool := &sync.Pool{
		New: func() interface{} {
			return "hello world"
		},
	}
	value := "hello xue yuan jun"
	pool.Put(value)
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}

func Pool1() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	pool := &sync.Pool{
		New: func() interface{} {
			return "hello world"
		},
	}
	go func(pool *sync.Pool, deferFunc func()) {
		defer func() {
			deferFunc()
		}()
		pool.Put("Hello xue yuan jun")
	}(pool, wg.Done)

	wg.Wait()
	fmt.Println(pool.Get())
}
