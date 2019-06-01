package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	//return map[string]int{"x": 1}
	strMap := make(map[string]int)
	for _, str := range strings.Fields(s) {
		//fmt.Println(str, strMap[str])
		_, ok := strMap[str]
		if ok {
			strMap[str] = strMap[str] + 1
		} else {
			strMap[str] = 1
		}
	}
	//fmt.Println(strMap)
	return strMap
}

func main() {
	wc.Test(WordCount)
}
