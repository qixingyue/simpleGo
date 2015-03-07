package main

import (
	"fmt"
	"learnGo/string_method/world"
	"reflect"
)

func main() {

	methodName := "A"
	h := new(world.Handler)
	t := reflect.TypeOf(h)
	v := reflect.ValueOf(h)
	if _, ok := t.MethodByName(methodName); !ok {
		fmt.Printf("No method found %s .... \n", methodName)
	} else {
		p := new(world.Param)
		m := v.MethodByName(methodName)
		params := []reflect.Value{reflect.ValueOf(p)}
		f := m.Call(params)
		l := len(f)
		fmt.Printf("%d\n", l)
		if f[0].Bool() {
			fmt.Printf("%s\n", f[1].String())
		}
	}

}
