// if文、if文の中でしか使わない変数をif文で定義できるという特徴がある
package main

import "fmt"

func main() {
	score := 61

	// 普通の条件分岐
	if score > 80 {
		fmt.Println("Great!")
	} else if score > 60 {
		fmt.Println("Good!")
	} else {
		fmt.Println("soso")
	}

	// if文の中で変数を定義する例
	if score := 52; score > 80 {
		fmt.Println("Great!")
	} else if score > 60 {
		fmt.Println("Good!")
	} else {
		fmt.Println("soso")
	}
}
