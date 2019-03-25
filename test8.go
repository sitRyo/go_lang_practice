// switch文
package main

import "fmt"

func main() {
	signal := "blue"

	switch signal {
	case "red":
		fmt.Println("Stop")
	case "yellow":
		fmt.Println("caution")
	case "green", "blue":
		fmt.Println("Go")
	default:
		fmt.Println("wrong")
	}

	score := 82
	switch {
	case score > 80:
		fmt.Println("Great!")
	default:
		fmt.Println("Bad")
	}
}
