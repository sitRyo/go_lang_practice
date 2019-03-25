// channel
package main

import (
	"fmt"
	"time"
)

func task1(result chan string) {
	time.Sleep(time.Second * 2)
	fmt.Println("task1 finished!")

	result <- "task1 result"
}

func task2() {
	fmt.Println("task2 finished!")
}

func main() {
	// chan型で受け渡すデータ型はstring
	result := make(chan string)

	// goをつけて並行処理にする
	go task1(result)
	go task2()

	// resultの中に何も入っていなければ、入ってくるまで待つ仕様になっている
	fmt.Println(<-result)

	// goroutuneが終わる前にmain関数が終わるため、待ち時間をつける
	time.Sleep(time.Second * 3)
}
