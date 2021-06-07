package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	start,end := 0,0
	for i := 0; i < len(s); i++ {
		pos := strings.Index(s[start:i], string(s[i]))
		if pos == -1{
			if i+1 > end{
				end = end + 1
			}
		} else {
			start = start + pos + 1
			end = end + pos + 1
		}
	}
	return end-start
}




func main() {
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
}


