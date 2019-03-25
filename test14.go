// 空インターフェースと型アサーション

package main

import "fmt"

type greeter interface {
	greet()
}

// 空インターフェースを引数にとると、どんな型の引数も受け取れるようになる
func show(t interface{}) {
	// ただし、空インターフェースでは型情報が欠落しているため、goでは動的に型を取得する方法を提供している
	// それを型アサーションといい、value, ok := <変数>.(<型>)で受け取る
	_, ok := t.(japanese)

	if ok {
		fmt.Println("i am japanese")
	} else {
		fmt.Println("i am not Japan")
	}

	// 型switch
	// t.(type)で型情報を取得することができる
	switch t.(type) {
	case japanese:
		fmt.Println("he is japanese")
	default:
		fmt.Println("he is not japanese")
	}
}

type japanese struct{}
type american struct{}

func (ja japanese) greet() {
	fmt.Println("こんにちは")
}

func (us american) greet() {
	fmt.Println("Hello")
}

// main function
func main() {
	greeters := []greeter{japanese{}, american{}}

	for _, greeter := range greeters {
		greeter.greet()
		show(greeter)
	}
}
