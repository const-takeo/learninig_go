package main

import (
	"fmt"
	"strings"
)

// 型の宣言が必須
func multiply(a, b int) int {
	return a * b
}

// return値を２個以上還せられる。
func lenAndUpper(name string) (int, string, string) {
	return len(name), strings.ToUpper(name), name
}

//複数の引数を受け取る方法
func repeatMe(words ...string) {
	fmt.Println(words)
}

// naked return & defer
func lenAndUpperTwo(name string) (length int, upper string) {
	defer fmt.Println("Defer is amazing")
	length = len(name)
	upper = strings.ToUpper(name)
	return
}

func main() {
	fmt.Println(multiply(2, 2))
	// ignore value
	len, upper, _ := lenAndUpper("Kunwoong")
	fmt.Println(len, upper)
	//arrayとして返還される
	repeatMe("yun", "kun", "woong")
	//
	length, uppperCase := lenAndUpperTwo("Kakao Talk")
	fmt.Println(length, uppperCase)
}
