package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	files, err := filepath.Glob("[")

	if err == nil {
		fmt.Println("matched files", files)
		return
	}

	if err == filepath.ErrBadPattern {
		fmt.Println("Bad pattern error:", err)
	} else {
		fmt.Println("Generic error:", err)
	}
}
