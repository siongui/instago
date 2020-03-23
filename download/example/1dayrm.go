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

func isOlderThanOneDay(t time.Time) bool {
	return time.Now().Sub(t) > 24*time.Hour
}

func processFile(path string, info os.FileInfo) {
	filename := filepath.Base(path)
	//if strings.Contains(filename, "-post-") || strings.Contains(filename, "-story") || strings.Contains(filename, "-postlive-") {
	if strings.Contains(filename, "-story") || strings.Contains(filename, "-postlive-") {
		//fmt.Println(filename)
		tmp := strings.Split(filename, "-")
		utimeext := tmp[len(tmp)-1]
		utime := strings.Split(utimeext, ".")[0]
		//fmt.Println(utime)

		i, err := strconv.ParseInt(utime, 10, 64)
		if err != nil {
			panic(err)
		}
		t := time.Unix(i, 0)
		//fmt.Println(t.Format(time.UnixDate))

		if isOlderThanOneDay(t) {
			fmt.Println(filename)
			fmt.Println(t.Format(time.UnixDate))
			//os.Remove(path)
		}
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
