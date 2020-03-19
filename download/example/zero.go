package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func processFile(path string, info os.FileInfo) {
	if info.Size() == 0 {
		fmt.Println(path, "=0")
	}
}

func main() {
	root := flag.String("root", "Instagram", "dir of downloaded files")
	flag.Parse()

	err := filepath.Walk(*root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.Mode().IsRegular() {
			processFile(path, info)
		}

		return nil
	})
	if err != nil {
		panic(err)
		return
	}
}
