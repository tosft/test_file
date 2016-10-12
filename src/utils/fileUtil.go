package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

/**
获取路径下的所有文件
*/
func GetFileList(path string) []string {
	var fileNames []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		fileNames = append(fileNames, path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return fileNames
}

func ReadFile(fileName string) []string {
	var lines []string
	f, err := os.Open(fileName) //打开文件
	defer f.Close()             //打开文件出错处理
	if nil == err {
		buff := bufio.NewReader(f) //读入缓存
		for {
			line, err := buff.ReadString('\n') //以'\n'为结束符读入一行
			if err != nil || io.EOF == err {
				break
			}
			lines = append(lines, line)
		}
	}
	return lines
}
