package main

import (
	"fmt"
)

func f1(m int) {
	fmt.Printf("%d\n", m)
}
func f1(m string) {
	fmt.Printf(m)
}
func main() {
	f1("Hello world.\n")
}
