package main

import (
	"fmt"
	_ "leafApi/models"
	"reflect"
	_ "strconv"
)

func main() {
	m := "hello world."
	fmt.Printf("%s\n", reflect.TypeOf(m))
	//fmt.Printf("%02d\n", 2)
	//m.HSet("world_config", "hello", strconv.Itoa(32))
	//if m.IsInSet("running_queue", "world", 3600) {
	//	fmt.Printf("IN")
	//} else {
	//	fmt.Printf("NOT IN")
	//}
}
