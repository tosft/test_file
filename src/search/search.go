package search

import (
	"strings"
)

const success = "success"

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
