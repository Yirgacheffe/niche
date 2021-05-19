package main

import (
	"fmt"
	"strings"
)

const smallwords = " a an on the to "

func main() {

	sampleStr := "this IS a Funny lOOking StRING!"

	fmt.Printf("Lower case: %s\n", strings.ToLower(sampleStr))
	fmt.Printf("Upper case: %s\n", strings.ToUpper(sampleStr))

	fmt.Printf("Title case: %s\n", strings.Title(strings.ToLower(sampleStr)))

	// Case #1
	properStr := properTitle("Welcome To The Dollhouse!")
	fmt.Printf("Proper case: %s\n", properStr)

	// Case #2
	betterStr := muchBetterTitle("a night to remember")
	fmt.Printf("Better case: %s\n", betterStr)

}

func properTitle(input string) string {

	words := strings.Fields(strings.ToLower(input))

	for idx, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[idx] = word
		} else {
			words[idx] = strings.Title(word)
		}
	}

	return strings.Join(words, " ")

}

func muchBetterTitle(input string) string {

	words := strings.Split(strings.ToLower(input), " ")

	for idx, word := range words {
		if strings.Contains(smallwords, " "+word+" ") && word != string(word[0]) {
			words[idx] = word
		} else {
			words[idx] = strings.Title(word)
		}
	}

	return strings.Join(words, " ")
}
