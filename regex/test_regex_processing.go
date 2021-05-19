package main

import (
	"fmt"
	"regexp"
)

func main() {

	a := "I am learning go language."
	re, _ := regexp.Compile("[a-z]{2,4}")

	first := re.Find([]byte(a))
	fmt.Println("Find: ", string(first))

	all := re.FindAll([]byte(a), -1)
	fmt.Printf("FindAll: %q\n", all)

	index := re.FindIndex([]byte(a))
	fmt.Printf("FindIndex: %q\n ", index)

	idxall := re.FindAllIndex([]byte(a), -1)
	fmt.Printf("FindAllIndex: %q\n", idxall)

	// case number 2
	re2, _ := regexp.Compile("am(.*)lang(.*)")

	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubMatch: ", submatch)

	for _, v := range submatch {
		fmt.Println(string(v))
	}

	submatchindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(submatchindex)

	submatchall := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println(submatchall)

	submatchallindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(submatchallindex)

}
