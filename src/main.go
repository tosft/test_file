package main

import (
	"flag"
	"fmt"
	"search"
	"utils"
)

func main() {
	flag.Parse()
	root := flag.Arg(0)
	fileNames := utils.GetFileList(root)
	for _, v := range fileNames {
		fmt.Println(v)
		lines := utils.ReadFile(v)
		containStrs := search.SliceContainsSuccess(lines)
		fmt.Println(len(containStrs))
		fmt.Println(containStrs[0])
	}
}
