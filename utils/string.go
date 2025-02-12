package utils

import (
	"strings"
)

func SubString(str string, start, num int) string {
	runes := []rune(str)
	strLen := len(runes)
	if start >= strLen {
		return ""
	}
	if num < 0 {
		return string(runes[start : strLen+num])
	}
	if start+num >= strLen || num == 0 {
		return string(runes[start:])
	}
	return string(runes[start : start+num])
}

// IfString 类似三目运算
func IfString(condition bool, str1 string, otherStr ...string) string {
	if condition {
		return str1
	}
	return StringOr(otherStr...)
}

// StringOr 尽量不返回空字符串
func StringOr(otherStr ...string) string {
	for _, str := range otherStr {
		if str != "" {
			return str
		}
	}
	return ""
}

// SnakeString 蛇形字符串
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

// CamelString 驼峰
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
