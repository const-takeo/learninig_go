package main

import "fmt"

type Person struct {
	name     string
	age      int
	favorite string
}

func main() {
	//structはmapと似ているがもっと柔軟
	//structを使う為にはtypeの宣言が必要
	//こういう書き方はよくないbecause上に宣言しておいたtypeを見ないと何が何か分からない。
	//goはclassがないalso　constructureもない
	yun := Person{"Yun", 30, "Kalbi"}
	fmt.Println(yun)
	//code wise
	kagawa := Person{name: "Kagawa", age: 30, favorite: "ramen"}
	fmt.Println(kagawa)
}
