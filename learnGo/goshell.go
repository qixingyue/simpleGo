package main

import(
	"fmt" 
	"os/exec"
	"strings"
	"bytes"
)

func main(){
	cmd := exec.Command("/data0/myget012/bin/mytget","-d","/tmp","-f","b.html","http://baidu.com")
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("EEEEEEE")
	}
	fmt.Printf("in all caps %s \n",out.String())
}
