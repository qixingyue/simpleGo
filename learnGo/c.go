package main

import (
	"fmt"
	"runtime"
)

func A() {
	panic(33)
}

func main() {
	defer func() {
		fmt.Println("\n+++++++++++++++++++++++++++++++++\n")
		fmt.Printf("%#v", Pos())
		funcName, file, line, ok := runtime.Caller(0)
		if ok {
			fmt.Println("Func Name=" + runtime.FuncForPC(funcName).Name())
			fmt.Printf("file: %s line=%d\n", file, line)
		}
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("\n+++++++++++++++++++++++++++++++++\n")
	}()
	A()
}
