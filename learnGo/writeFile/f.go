package main

import (
	"fmt"
	"math/rand"
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

func childReport() {
	for m := 0; m < 1000; m++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		diskHash := r.Intn(20)
		t := fmt.Sprintf("%d.txt", diskHash)
		text := fmt.Sprintf("%d\t%s\n", diskHash, t)
		fmt.Printf(text)
		ReportFile("a.txt", text)
	}
}

func main() {
	for i := 0; i < 5; i++ {
		go childReport()
	}
	var x string
	fmt.Scanln("%s", &x)
}
