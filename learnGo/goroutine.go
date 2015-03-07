package main 

import (
	"fmt"
	"runtime"
)

func async(s string) {
	for i := 0; i < 100; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main(){
	go async("world")
	async("Hello")
}
