package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var count int = 0

func isOlderThanOneDay(t time.Time) bool {
	return time.Now().Sub(t) > 24*time.Hour
}

func getTimeFromFilename(filename string) (t time.Time) {
	//fmt.Println(filename)
	tmp := strings.Split(filename, "-")
	utimeext := tmp[len(tmp)-1]
	utime := strings.Split(utimeext, ".")[0]
	//fmt.Println(utime)

	i, err := strconv.ParseInt(utime, 10, 64)
	if err != nil {
		panic(err)
	}
	t = time.Unix(i, 0)
	//fmt.Println(t.Format(time.UnixDate))
	return
}

func moveFile(todir, path string, info os.FileInfo) {
	count++
	fmt.Print(count, ": ")

	//fmt.Println(info.Name())
	//return

	newpath := filepath.Join(todir, info.Name())
	fmt.Printf("move %q to %q\n", path, newpath)
	err := os.Rename(path, newpath)
	if err != nil {
		panic(err)
	}
}

func processFile(todir, path string, info os.FileInfo) {
	// move files older than one day in name
	if isOlderThanOneDay(getTimeFromFilename(info.Name())) {
		moveFile(todir, path, info)
		return
	}

	// move files with "-post-" in the filename
	if strings.Contains(info.Name(), "-post-") {
		moveFile(todir, path, info)
		return
	}

	// move files older than one day
	if isOlderThanOneDay(info.ModTime()) {
		moveFile(todir, path, info)
		return
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
