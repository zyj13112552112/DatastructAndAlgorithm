package 中等

import "strings"

func strStr(haystack string, needle string) int {
	contains := strings.Contains(haystack, needle)
	if contains==false {return -1}
	length:=len(needle)
	for i:=0;i<=len(haystack)-length;i++{
		if haystack[i:i+length] == needle {
			return i
		}
	}
	return 1
}
