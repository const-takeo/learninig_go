package main

import "fmt"

func main() {
	//go's arrayの宣言
	//長さを特定する必要がある。
	names := [3]string{"Yun", "Kun", "Woong"}
	fmt.Println(names)

	//slice　is ...
	//長さが特定されていないarray
	lastnames := []string{"Yun", "Kim", "Park", "Yamashita"}
	fmt.Println(lastnames)
	//sliceに値を追加する為には append()関数を使用する。
	//append()関数は直接sliceを修正しない(It doesn't modify)
	lastnames = append(lastnames, "Nagasawa", "Kagawa")
	fmt.Println(lastnames)
}
