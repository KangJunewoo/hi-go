package practice

import (
	"fmt"

	"github.com/KangJunewoo/hi-go/mydict"
)

func dictionary() {
	// dictionary := mydict.Dictionary{"hello": "world"}
	// fmt.Println(dictionary)
	// definition, err := dictionary.Search("second")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}
