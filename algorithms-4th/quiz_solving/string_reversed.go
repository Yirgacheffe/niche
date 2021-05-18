package main

import "fmt"

func reverseNotAffectSpecial(s string) string {

	isNum := func(a byte) bool {
		return '0' <= a && a <= '9'
	}

	isAlphabet := func(a byte) bool {
		return ('a' <= a && a <= 'z') || ('A' <= a && a <= 'Z')
	}

	chars := []byte(s)

	// var temp []byte
	temp := []byte{}

	// Create temp not include special char
	for _, v := range chars {

		if isNum(v) || isAlphabet(v) {
			temp = append(temp, v)
		}
	}

	rs := reverseString(string(temp))

	// Loop input and reversed array, set value back
	j := 0
	for i, v := range chars {

		if isNum(v) || isAlphabet(v) {
			chars[i] = rs[j]
			j++
		}

	}

	return string(chars)

}

func reverseString(s string) string {

	chars := []byte(s)

	lo := 0
	hi := len(chars) - 1

	for lo < hi {
		t := chars[lo]
		chars[lo] = chars[hi]
		chars[hi] = t

		lo++
		hi--

	}

	return string(chars)

}

func main() {

	s := "Ab,c,de!$4"

	fmt.Println(s)
	fmt.Println(reverseNotAffectSpecial(s))

}
