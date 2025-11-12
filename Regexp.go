package main

import (
	"fmt"
	"regexp"
)

//func main() {
//	replace()
//}

func match() {
	pattern := `^[a-zA-Z0-9]+$`
	regex := regexp.MustCompile(pattern)

	str := "Hello123"
	if regex.MatchString(str) {
		fmt.Println("字符串匹配正则表达式")
	} else {
		fmt.Println("字符串不匹配正则表达式")
	}
}

func find() {
	pattern := `\d+`
	regex := regexp.MustCompile(pattern)

	str := "我有 3 个苹果和 5 个香蕉"
	matches := regex.FindAllString(str, -1)
	fmt.Println("找到的数字：", matches)
}

func replace() {
	pattern := `\s+`
	regex := regexp.MustCompile(pattern)

	str := "Hello    World"
	result := regex.ReplaceAllString(str, " ")
	fmt.Println("替换后的字符串：", result)
}
