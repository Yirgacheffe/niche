package main

import "fmt"

func isPalindrome(s string) bool {

	lo := 0
	hi := len(s) - 1

	for lo <= hi {
		if s[lo] == s[hi] {
			lo++
			hi--
		} else {
			return false
		}
	}

	return true

}

func substring(s string, b, e int) string {
	rs := []rune(s)
	return string(rs[b : e+1])
}

func main() {

	s := "abaaabaaaba"

	s1 := substring(s, 0, 10)
	s2 := substring(s, 2, 5)
	s3 := substring(s, 5, 9)
	s4 := substring(s, 2, 3)

	fmt.Printf("%s is palindrome: %v\n", s1, isPalindrome(s1))
	fmt.Printf("%s is palindrome: %v\n", s2, isPalindrome(s2))
	fmt.Printf("%s is palindrome: %v\n", s3, isPalindrome(s3))
	fmt.Printf("%s is palindrome: %v\n", s4, isPalindrome(s4))

}
