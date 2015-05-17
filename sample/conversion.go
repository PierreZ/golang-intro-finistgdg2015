package main

import (
	"fmt"
	"reflect"
)

func main() {

	i := 42
	fmt.Println(i, reflect.TypeOf(i))
	f := float64(i)
	fmt.Println(f, reflect.TypeOf(f))
	u := uint(f)
	fmt.Println(u, reflect.TypeOf(u))
}
