package something

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

// SayGoodbye 関数の頭文字が大文字の場合export出来る。
func SayGoodbye() {
	fmt.Println("Goodbye")
}
