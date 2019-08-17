package utils

import (
	"math/rand"
	"regexp"
	"time"
)

// GetRandomString get random string from "0123456789abcdefghijklmnopqrstuvwxyz"
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// StrReplaceAll Replace string with pattern
func StrReplaceAll(s string, pattern string, repl string) string {
	if ok, _ := regexp.Match(pattern, []byte(s)); ok {
		re, _ := regexp.Compile(pattern)
		final := re.ReplaceAllString(s, repl)
		return final
	}
	return s
}

// StrDetect Detect string pattern
func StrDetect(s string, pattern string) bool {
	if ok, _ := regexp.Match(pattern, []byte(s)); ok {
		return true
	}
	return false
}

// StrSplit Split string with pattern
func StrSplit(s string, pattern string, n int) (arr []string) {
	if ok, _ := regexp.Match(pattern, []byte(s)); ok {
		re, _ := regexp.Compile(pattern)
		arr := re.Split(s, n)
		return arr
	}
	return []string{s}
}

// StrExtract extract string with pattern
func StrExtract(s string, pattern string, n int) (arr string) {
	if ok, _ := regexp.Match(pattern, []byte(s)); ok {
		re, _ := regexp.Compile(pattern)
		arr := re.FindString(s)
		return arr
	}
	return s
}
