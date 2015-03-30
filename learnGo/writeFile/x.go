package main

import (
	"fmt"
	_ "math/rand"
	"os"
	"time"
)

func ReportFile(fileName string, text string) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0x644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(text)
}

func f(n int) {
	for i := 0; i < 10000; i++ {
		text := fmt.Sprintf("%d\t%d\t%d\n", n, i, i)
		ReportFile("a.txt", text)
		//r := rand.New(rand.NewSource(time.Now().UnixNano()))
		time.Sleep(200)
	}
}
func main() {
	for index := 0; index < 10; index++ {
		go f(index)
	}
	var input string
	fmt.Scanln(&input)
}
