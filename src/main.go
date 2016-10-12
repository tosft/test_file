package main

import (
	"flag"
	"fmt"
	"net/http"
	"search"
	"strings"
	"utils"
)

var path string
var fileName string

func main() {
	flag.Parse()
	arg := flag.Arg(0)
	strs := strings.Split(arg, ",")
	path = strs[0]
	fileName = strs[1]

	fmt.Println("start 18080....")
	http.HandleFunc("/", register)
	http.ListenAndServe(":18080", nil)

}

func register(w http.ResponseWriter, r *http.Request) {
	fileNames := utils.GetFileList(path)
	var strs []string
	for _, v := range fileNames {

		if strings.Contains(v, fileName) {
			size, sum := search.GetFileSuccessSizeAndSum(v)
			strs = append(strs, search.SpliceStr(v, size, sum))
		}
	}
	fmt.Fprintln(w, strsToString(strs))
}

func strsToString(strs []string) string {

	var str string
	str += "<div>"
	for _, v := range strs {
		str += v + " <br/>"
	}
	str += "</div>"
	return str
}
