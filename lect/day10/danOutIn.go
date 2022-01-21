package main

import "fmt"

func producer(first chan int) {
	defer close(first) // channelを閉じる
	for i := 0; i < 10; i++ {
		first <- i
	}
}

func multi2(first, second chan int) {
	defer close(second)
	for i := range first {
		second <- i * 2
	}
}

func multi4(second, third chan int) {
	defer close(third)
	for i := range second {
		third <- i * 2
	}
}

func main() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producer(first)
	go multi2(first, second)
	go multi4(second, third)

	for res := range third {
		fmt.Println(res)
	}
}

// goroutin同士の連携
