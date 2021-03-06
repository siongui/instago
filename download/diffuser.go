package igdl

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/siongui/instago"
)

// Set Difference: users1 - users0
func DiffFollowUsers(users0, users1 []instago.IGFollowUser) {
	oriNames := make(map[int64]bool)
	for _, user := range users0 {
		oriNames[user.Pk] = true
	}

	for _, user := range users1 {
		if _, ok := oriNames[user.Pk]; !ok {
			log.Println("https://www.instagram.com/" + user.Username + "/")
			log.Println(user)
		}
	}
}

// Given directory which contains data saved by running at leat twice
// SaveSelfFollow(), find out who blocks, disappears, kicks off. Or newly added.
func DiffFollowData(dir, keyword string) (err error) {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	filenames := []string{}
	for _, info := range infos {
		if strings.Contains(info.Name(), keyword) {
			filenames = append(filenames, info.Name())
		}
	}

	if len(filenames) < 2 {
		log.Println("no enough data to diff")
		return
	}

	sort.Strings(filenames)

	x0 := filenames[len(filenames)-2]
	x1 := filenames[len(filenames)-1]

	users0, err := LoadFollowUsers(path.Join(dir, x0))
	if err != nil {
		return
	}
	users1, err := LoadFollowUsers(path.Join(dir, x1))
	if err != nil {
		return
	}

	log.Println("New or re-appear")
	log.Println(x1, "-", x0)
	DiffFollowUsers(users0, users1)
	log.Println("Blocked, Kicked off, or Disappear")
	log.Println(x0, "-", x1)
	DiffFollowUsers(users1, users0)
	return
}

func GetLatestFile(dir, keyword string) (latest os.FileInfo, err error) {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	if len(infos) == 0 {
		err = errors.New("no files in dir")
		return
	}

	latest = infos[0]
	for _, info := range infos {
		if strings.Contains(info.Name(), keyword) {
			if info.ModTime().After(latest.ModTime()) {
				latest = info
			}
		}
	}

	if !strings.Contains(latest.Name(), keyword) {
		err = errors.New("no filename contains keyword")
		return
	}

	return
}

func LoadLatestFollowingUsers() (users []instago.IGFollowUser, err error) {
	latestfile, err := GetLatestFile(GetFollowDir(), "-following-")
	if err != nil {
		return
	}

	p := path.Join(GetFollowDir(), latestfile.Name())
	return LoadFollowUsers(p)
}
