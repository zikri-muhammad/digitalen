// hands on tugas 3

package main

import (
	"fmt"
	"reflect"
)

func main() {

	// type data

	i := 42
	f := 3.143
	g := 0.867 + 5i

	// print tyoe data
	fmt.Println("type data", i, "adalah", reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(g))
}
