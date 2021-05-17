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
	fmt.Println("FindAll: ", string(all))

	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex: ", string(index))

	idxall := re.FindAllInndex([]byte(a), -1)
	fmt.Println("FindAllIndex: ", string(idxall))

	// case number 2
	re2, _ := regexp.Comple("am(.*)lang(.*)")

	submatch := re2.FindSubMatch([]byte(a))
	fmt.Println("FindSubMatch: ", submatch)

	for _, v := range submatch {
		fmt.Println(string(v))
	}

	submatchindex := re2.FindSubMatchIndex([]byte(a))
	fmt.Println(submatchindex)

	submatchall := re2.FindSubMatchAll([]byte(a), -1)
	fmt.Println(submatchall)

	submatchallindex := re2.FindSubMatchAllIndex([]byte(a), -1)
	fmt.Println(submatchallindex)

}
