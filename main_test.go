package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	input := "racecar"
	if !isPalindrome(input) {
		t.Errorf("Error! String %v is a palindrome", input)
	}
}

func TestCase2(t *testing.T) {
	input := "golang"
	if isPalindrome(input) {
		t.Errorf("Error! String %v is not a palindrome", input)
	}
}
