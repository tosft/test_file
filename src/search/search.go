package search

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"utils"
)

const (
	success = "success"
	price   = "price:"
)

func SliceContainsSuccess(strs []string) []string {
	var containStrs []string
	for _, v := range strs {
		if strContainSuccess(v) {
			containStrs = append(containStrs, v)
		}
	}
	return containStrs
}

func strContainSuccess(str string) bool {
	return strings.Contains(str, success)
}

func strContainPrice(str string) bool {
	return strings.Contains(str, price)
}

func SplitStr(str string) []string {
	return strings.Split(str, " ")
}

func replacePrice(str string) int {
	s := strings.Replace(str, price, "", -1)
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}
	return num
}

func Sum(str string) int {
	var sum int
	for _, v := range SplitStr(str) {
		if strContainPrice(v) {
			sum += replacePrice(v)
		}
	}
	return sum
}

func SumFile(lines []string) int {
	var sum int
	for _, v := range lines {
		sum += Sum(v)
	}
	return sum
}

func SpliceStr(fileName string, size int, sum int) string {
	splitStrs := strings.Split(fileName, "_")
	date := strings.Split(splitStrs[1], ".")[0]
	str := fmt.Sprintf("日期:%s  充值笔数:%s  总金额:%s", date, strconv.Itoa(size), strconv.Itoa(sum))
	return str
}

func GetFileSuccessSizeAndSum(fileName string) (size, sum int) {
	lines := utils.ReadFile(fileName)
	containStrs := SliceContainsSuccess(lines)
	sum = SumFile(containStrs)
	size = len(containStrs)
	return
}
