package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func processDir(path string, info os.FileInfo) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	if len(files) != 0 {
		return
	}

	err = os.Remove(path)
	if err != nil {
		panic(err)
	}

	fmt.Println(path, "removed!")
}

func main() {
	root := flag.String("root", "Instagram", "dir of downloaded files")
	flag.Parse()

	err := filepath.Walk(*root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			processDir(path, info)
		}

		return nil
	})
	if err != nil {
		panic(err)
		return
	}
}
