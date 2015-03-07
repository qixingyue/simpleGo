package main 

import (
	"fmt" 
	"runtime"
)

func sum(a []int, c chan int) {
	total := 0
	for _,v := range a {
		total += v 
	}
	c <- total
}

func main(){
	runtime.GOMAXPROCS(8) 	
	a := []int{1,2,3,4,5,6,7,8}
	c := make(chan int)
	go sum(a[:5],c)
	go sum(a[5:],c)
	x,y := <-c, <-c
	fmt.Println(x,y,x+y)

}
