package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDir(wg *sync.WaitGroup, dir string, fileSize chan<- int64) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(wg, subdir, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

// sema is counting semaphore to limie the concurrency in dir enters
var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

var verbose = flag.Bool("v", false, "show progress verbose message")

func main() {

	// Determine the initial directories
	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	filesize := make(chan int64)
	var nfiles, nbytes int64

	var wg sync.WaitGroup

	for _, dir := range roots {
		wg.Add(1)
		go walkDir(&wg, dir, filesize)
	}

	go func() {
		wg.Wait()
		close(filesize)
	}()

loop:
	for {
		select {
		case size, ok := <-filesize:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}
