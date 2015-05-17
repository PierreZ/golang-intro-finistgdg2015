package main

import (
	"fmt"
	"reflect"
)

func main() {

	var i int // Java style
	i = 10

	foo := "bar" // Go style

	fmt.Println(i, reflect.TypeOf(i))
	fmt.Println(foo, reflect.TypeOf(foo))
}
