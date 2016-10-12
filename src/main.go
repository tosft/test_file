package main

import (
	"flag"
	"fmt"
	"search"
	"strconv"
	"strings"
	"utils"
)

func main() {
	flag.Parse()
	root := flag.Arg(0)
	fileNames := utils.GetFileList(root)

	for _, v := range fileNames {
		size, sum := test(v)
		fmt.Println(spiltStr(v, size, sum))
	}

}

func spiltStr(fileName string, size int, sum int) string {
	splitStrs := strings.Split(fileName, "_")
	date := strings.Split(splitStrs[1], ".")[0]
	str := fmt.Sprintf("日期:%s  充值笔数:%s  总金额:%s", date, strconv.Itoa(size), strconv.Itoa(sum))
	return str
}

func test(fileName string) (size, sum int) {
	lines := utils.ReadFile(fileName)
	containStrs := search.SliceContainsSuccess(lines)
	sum = search.SumFile(containStrs)
	size = len(containStrs)
	return
}
