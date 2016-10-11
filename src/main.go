package main

import (
	"flag"
	"utils"
)

func main() {
	flag.Parse()
	root := flag.Arg(0)
	util.GetFileList(root)
}
