package main

import "fmt"

func main() {
	// 一般の書き方
	const name string = "Yun"
	// コントラクション方
	// 関数内でしか使え無い。
	secondName := "Yun"
	fmt.Println(secondName)
}
