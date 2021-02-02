package main

import "fmt"

func main() {
	//goのmapはjsやpythonのobjectと似ているが違う
	carsPowers := map[string]int{"WRX": 300, "S660": 64, "MX-5": 140}
	fmt.Println(carsPowers)
	//for文も使える。(iterable)
	for key, value := range carsPowers {
		fmt.Println(key, value)
	}
	//mapに値を追加する。
}
