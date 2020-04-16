package igdl

import (
	"fmt"
	"io/ioutil"
)

func saveIdUsername(id, username string) (err error) {
	p := getIdUsernamePath(id, username)
	CreateFilepathDirIfNotExist(p)
	// check if file exist
	if _, err := os.Stat(p); os.IsNotExist(err) {
		// file not exists
		err = ioutil.WriteFile(p, []byte(""), 0644)
		if err == nil {
			fmt.Println(id, username, "saved")
		}
	}
	return
}
