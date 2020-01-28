package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func isOlderThanOneDay(t time.Time) bool {
	return time.Now().Sub(t) > 24*time.Hour
}

func processFile(todir, path string, info os.FileInfo) {
	if isOlderThanOneDay(info.ModTime()) {
		newpath := filepath.Join(todir, info.Name())
		fmt.Printf("move %q to %q\n", path, newpath)
		err := os.Rename(path, newpath)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	root := flag.String("root", "Instagram", "dir of downloaded files")
	todir := flag.String("todir", "~/", "dir to which files are moved")
	flag.Parse()

	err := filepath.Walk(*root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.Mode().IsRegular() {
			processFile(*todir, path, info)
		}

		return nil
	})
	if err != nil {
		panic(err)
		return
	}
}
