package main

import (
	"fmt"

	"github.com/const-takeo/learning_go/dict/dictionary"
)

func main() {
	mydict := dictionary.Dictionary{"first": "FirstWord"}
	// mydict["hello"] = "hello"
	fmt.Println(mydict["first"])
	//use method
	val, err := mydict.Search("first")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
	//add
	addErr := mydict.Add("Hello", "Greeting")
	if addErr != nil {
		fmt.Println(addErr)
	}
	fmt.Println(mydict["Hello"])
}
