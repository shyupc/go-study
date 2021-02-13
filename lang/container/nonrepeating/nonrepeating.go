package main

import "fmt"

func lengthNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(lengthNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthNonRepeatingSubStr("bbbbbb"))
	fmt.Println(lengthNonRepeatingSubStr("pwwkew"))

	fmt.Println(lengthNonRepeatingSubStr(""))
	fmt.Println(lengthNonRepeatingSubStr("b"))
	fmt.Println(lengthNonRepeatingSubStr("abcdef"))

	fmt.Println(lengthNonRepeatingSubStr("这里是慕课网"))
	fmt.Println(lengthNonRepeatingSubStr("一二三二一"))
	fmt.Println(lengthNonRepeatingSubStr("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
