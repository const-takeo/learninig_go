package main

import (
	"fmt"

	"github.com/const-takeo/learning_go/condition/switchTest"
)

func canIDrink(age int) bool {
	//goではif文の中に変数の宣言ができる。
	//variable expression
	if koreanAge := age + 2; koreanAge >= 20 {
		return true
	} //if blockがreturnで終わるとelse blockを省略できる。
	return false
}

func main() {
	result := canIDrink(18)
	fmt.Println(result)
	judge := switchTest.CanIDrinkTwo(51)
	fmt.Println(judge)
}
