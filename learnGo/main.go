package main

import(
	"fmt"
)

type testFun func(int) 

func ra(b int){
	fmt.Printf("Dynamic code %d\n",b)
}


func doA(b testFun){
	m := 128 
	b(m)	
}

func add1( a *int) int {
	defer fmt.Printf("1")
	defer fmt.Printf("2")
	defer fmt.Printf("3")
	defer fmt.Printf("4")
	*a += 1
	return *a
}

func main(){
	intVar := 30
	var a bool
	b := false
	add1(&intVar)	
	fmt.Printf("%v %v  \n",a,b)
	fmt.Printf("Should Be 31 , \n %v   \n",intVar)
	doA(ra)
}
