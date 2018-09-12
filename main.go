package main

import "fmt"

func main() {
	str := "racecar"
	fmt.Println(isPalindrome(str))
}

func isPalindrome(str string) bool {
	left := 0
	right := len(str) - 1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
