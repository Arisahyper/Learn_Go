package main

import "fmt"

func main() {
	ch := make(chan int, 2) // チャネルを用意
	ch <- 100               // チャネルにデータを送る
	fmt.Println(len(ch))    // チャネルの要素数を表示 / 1
	ch <- 200               // チャネルにデータを送る / もう入らない
	fmt.Println(len(ch))    // チャネルの要素数を表示 / 2

	x := <-ch // チャネルからxへデータを受け取る / 一回取り出すことによって要素数が減る
	fmt.Println(x)
	ch <- 300            // チャネルにデータを送る
	fmt.Println(len(ch)) // チャネルの要素数を表示 / 2
}
