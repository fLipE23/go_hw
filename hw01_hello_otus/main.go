package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {

	baseString := "Hello, OTUS!"
	resultString := stringutil.Reverse(baseString)

	fmt.Println(resultString)

}
