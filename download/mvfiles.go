package igdl

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func IsOlderThanOneDay(t time.Time) bool {
	return time.Now().Sub(t) > 24*time.Hour
}

func GetTimeFromStoryFilename(filename string) (t time.Time, err error) {
	if !strings.Contains(filename, "-story-") {
		err = errors.New("invalid file name: " + filename)
		return
	}

	tmp := strings.Split(filename, "-")
	if len(tmp) == 0 {
		err = errors.New("invalid file name: " + filename)
		return
	}

	utimeext := tmp[len(tmp)-1]
	tmp2 := strings.Split(utimeext, ".")
	if len(tmp2) != 2 {
		err = errors.New("invalid file name: " + filename)
		return
	}
	utimestr := tmp2[0]

	i, err := strconv.ParseInt(utimestr, 10, 64)
	if err != nil {
		return
	}
	t = time.Unix(i, 0)
	//println(t.Format(time.UnixDate))
	return
}

func MoveFileToDir(todir, path string, info os.FileInfo) error {
	newpath := filepath.Join(todir, info.Name())
	return os.Rename(path, newpath)
}

func moveStoryFile(todir, path string, info os.FileInfo) (err error) {
	// move story file with modification time older than one day
	if IsOlderThanOneDay(info.ModTime()) {
		return MoveFileToDir(todir, path, info)
	}

	// move story file name older than one day
	t, err := GetTimeFromStoryFilename(info.Name())
	if err != nil {
		return
	}
	if IsOlderThanOneDay(t) {
		return MoveFileToDir(todir, path, info)
	}

	return
}

func MoveExpiredStory(storydir, todir string) (err error) {
	err = filepath.Walk(storydir, func(path string, info os.FileInfo, e error) error {
		if e != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, e)
			return e
		}
		if info.Mode().IsRegular() {
			if strings.Contains(info.Name(), "-story-") {
				e2 := moveStoryFile(todir, path, info)
				if e2 != nil {
					return e2
				}
			}
		}

		return nil
	})
	return
}
