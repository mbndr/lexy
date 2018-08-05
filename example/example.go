// only for lexing purpose
package main

import "fmt"

var x = 42

/*
This is a multi line comment
*/
func HelloWorld() {
	fmt.Println("Hell\"o World!")

	_ = len([]string{"test"})
}
