package main

import "fmt"

// goはループ処理の時forだけで行う。
func superAdd(numbers ...int) int {
	defer fmt.Println("Go is awesome")
	total := 0
	//rangeを用いてforの作成。
	for index, count := range numbers {
		fmt.Println(index, count)
	}
	//
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(result)
}
