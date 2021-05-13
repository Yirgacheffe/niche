package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Don't run this, this file is a code backup
// Not ready to run......
func copyFS(src, dst string) {

	os.MkdirAll(dst, 0777)
	srcFileStat, _ := os.Lstat(src)

	if isDir := srcFileStat.IsDir(); isDir {
		srcDirContent, _ := ioutil.ReadDir(src)

		for _, contentFile := range srcDirContent {
			srcFilePath := filepath.Join(src, contentFile.Name())
			dstDir := filepath.Join(dst, srcFileStat.Name())

			copyFS(srcFilePath, dstDir)
		}

	} else {
		srcContent, _ := ioutil.ReadFile(src)
		dstFilePath := filepath.Join(dst, srcFileStat.Name)

		ioutil.WriteFile(dstFilePath, srcContent, 0777)
		fmt.Println(dstFilePath, "->", dstFilePath)
	}

}

func main() {
	copyFS("sub/nested", "my/dir")
}
