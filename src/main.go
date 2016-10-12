package main

import (
	"flag"
	"fmt"
	"strings"
	"utils"
)

const success = "success"

func main() {
	flag.Parse()
	root := flag.Arg(0)
	fileNames := utils.GetFileList(root)
	for _, v := range fileNames {
		fmt.Println(v)
		lines := utils.ReadFile(v)
		num := sliceContainsSuccess(lines)
		fmt.Println(num)
	}
}

func sliceContainsSuccess(strs []string) int32 {
	var num int32 = 0
	for _, v := range strs {
		if strContainSuccess(v) {
			num++
		}
	}
	return num
}

func strContainSuccess(str string) bool {
	return strings.Contains(str, success)
}
