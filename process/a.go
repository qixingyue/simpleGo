package main

import (
	"fmt"
	"os"
)

func customSignal(c chan os.Signal) {
	s := <-c
	fmt.Sprintf("%#v", s)
}

func main() {
	c := make(chan os.Signal, 1)
	go customSignal(c)
	fmt.Printf("ARGC:%d\n", len(os.Args))
	fmt.Printf("Process Name: %s \n", os.Args[0])
	fmt.Printf("PID : %d\n", os.Getpid())
	fmt.Printf("PPID : %d\n", os.Getppid())
	for {
	}
}
