package main

import(
	"fmt"
	"os/exec"
	"bytes"
	"strings"
)

func checkMd5(fileName string, aimMd5 string) bool {
	cmd := exec.Command("md5sum",fileName)
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%#v",err)
	}
	resultString := string(out.Bytes())
	resultArr := strings.Split(resultString," ")
	return strings.EqualFold(resultArr[0], aimMd5)
}

func main(){
	fileName := "t.go"
	aimMd5 := "57c054f64a7506f94dc4be5f5027255b"
	if checkMd5(fileName,aimMd5) {
		fmt.Printf("OK")
	}
}
