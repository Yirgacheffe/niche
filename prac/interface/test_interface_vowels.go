package main

import "fmt"

type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rv := range ms {
		if rv == 'a' || rv == 'e' || rv == 'i' || rv == 'o' || rv == 'u' {
			vowels = append(vowels, rv)
		}
	}
	return vowels
}

func main() {
	var name = MyString("Gem DengZiQi")
	var v VowelsFinder

	v = name
	fmt.Printf("Vowels are %c", v.FindVowels())
}
