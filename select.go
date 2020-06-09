package main

import (
	"fmt"
	"math/rand"
)

func ISelect(){
	chs := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	index:= rand.Intn(3)
	fmt.Printf("随机索引/通道：%d\n", index)

	chs[index] <- index

	select {
		case <- chs[0]:
			fmt.Println("zero")
		case <- chs[1]:
			fmt.Println("one")
		case <- chs[2]:
			fmt.Println("two")
		default:
			fmt.Println("default")
	}
}
