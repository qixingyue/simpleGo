package tools

import (
	"fmt"
)

func init(){
	fmt.Printf("Hello import.")
}

type U struct {
	Url	string
}

func (this *U) Hello() {
	fmt.Printf("Hello wrold.")
}

func GetU() *U {
	return new(U)
}

func HelloPackage() {
	fmt.Printf("Hello package .")
}
