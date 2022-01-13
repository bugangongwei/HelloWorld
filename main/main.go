package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

func main() {
	// net.TcpTest()

	_, filename, _, _ := runtime.Caller(1)
	dir := path.Dir(filename)

	content, err := os.ReadFile(dir + "/HelloWorld/verification/my.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}