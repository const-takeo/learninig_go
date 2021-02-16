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
	_, err := mydict.Search("first")
	if err != nil {
		fmt.Println(err)
	}
	//add
	addErr := mydict.Add("Hello", "Greeting")
	if addErr != nil {
		fmt.Println(addErr)
	}
	//iterable
	for key, value := range mydict {
		fmt.Println(key, value)
	}
	//Update
	updateErr := mydict.Update("first", "こんばんは")
	if updateErr != nil {
		fmt.Println((updateErr))
	}
	fmt.Println(mydict.Search("first"))

	//Delete
	delErr := mydict.Delete("first")
	if delErr != nil {
		fmt.Println(delErr)
	}
	_, confirmerr := mydict.Search("first")
	fmt.Println(confirmerr)
}
