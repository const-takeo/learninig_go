package main

import "fmt"

func main() {
	a := 2
	//この時点でaの値がコピーされる。
	b := &a
	a = 10
	// * <- see through , look through, アドレス住所が乗っているところで使える。
	fmt.Println(a, *b)

	// 反対に*付のbを用いてaの値を直接書き換える事も可能。
	// because b is pointer of a
	*b = 20
	fmt.Println(a)
}
