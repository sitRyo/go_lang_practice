package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		} else if i == 8 {
			break
		}
		fmt.Println(i)
	}

	n := 0
	for n < 10 {
		fmt.Println(n)
		n++
	}
}
